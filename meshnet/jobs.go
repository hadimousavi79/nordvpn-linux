package meshnet

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/NordSecurity/nordvpn-linux/events"
	"github.com/NordSecurity/nordvpn-linux/internal"
	meshInternal "github.com/NordSecurity/nordvpn-linux/meshnet/internal"
	"github.com/NordSecurity/nordvpn-linux/meshnet/jobs"
)

const ()

func (s *Server) StartJobs(
	meshnetPublisher events.PublishSubcriber[bool],
) {
	if _, err := s.scheduler.NewJob(gocron.DurationJob(2*time.Hour), gocron.NewTask(JobRefreshMeshnet(s)), gocron.WithName("job refresh meshnet")); err != nil {
		log.Println(internal.WarningPrefix, "job refresh meshnet schedule error:", err)
	}

	if _, err := s.scheduler.NewJob(
		gocron.DurationJob(1*time.Second),
		gocron.NewTask(JobMonitorFileshareProcess(s)),
		gocron.WithName("job monitor fileshare process")); err != nil {
		log.Println(internal.WarningPrefix, "job monitor fileshare process schedule error:", err)
	}

	s.scheduler.Start()
	for _, job := range s.scheduler.Jobs() {
		err := job.RunNow()
		if err != nil {
			log.Println(internal.WarningPrefix, job.Name(), "first run error:", err)
		}
	}

	// monitors the meshnet status and starts/stops the peers refreshing job
	meshnetPublisher.Subscribe(func(enabled bool) error {
		// TODO: check what happens if meshnet is started
		if enabled {
			job, err := s.scheduler.NewJob(
				gocron.DurationRandomJob(internal.MeshnetPeersUpdateInterval/2, internal.MeshnetPeersUpdateInterval),
				gocron.NewTask(jobs.JobUpdateMeshnetPeers(s, s, s.dataManager, s.subjectPeerUpdate)),
				gocron.WithName("refresh peers"),
				gocron.WithTags(internal.MeshnetPeersJobTag),
			)
			if err != nil {
				log.Println(internal.ErrorPrefix, "failed to schedule peers refresh job", err)
				return err
			}

			log.Println(internal.DebugPrefix, "peers refresh job scheduled")

			if err := job.RunNow(); err != nil {
				log.Println(internal.ErrorPrefix, "failed to run peers refresh job", err)
				return err
			}
		} else {
			s.scheduler.RemoveByTags(internal.MeshnetPeersJobTag)
			log.Println(internal.DebugPrefix, "stop peer refresh job")
		}
		return nil
	})
}

func JobRefreshMeshnet(s *Server) func() error {
	return func() error {
		// ignore what is returned, try to do it here as light as possible
		_, _ = s.RefreshMeshnet(nil, nil)
		return nil
	}
}

func JobMonitorFileshareProcess(s *Server) func() error {
	job := monitorFileshareProcessJob{
		isFileshareAllowed: false,
		meshChecker:        s,
		rulesController:    s.netw,
		processChecker:     defaultProcessChecker{},
	}
	return job.run
}

func (j *monitorFileshareProcessJob) run() error {
	if !j.meshChecker.IsMeshnetOn() {
		if j.isFileshareAllowed {
			if err := j.rulesController.ForbidFileshare(); err == nil {
				j.isFileshareAllowed = false
			}
		}
		return nil
	}

	if j.processChecker.isFileshareRunning() {
		j.rulesController.PermitFileshare()
		j.isFileshareAllowed = true
	} else {
		j.rulesController.ForbidFileshare()
		j.isFileshareAllowed = false
	}

	return nil
}

type defaultProcessChecker struct{}

func (defaultProcessChecker) isFileshareRunning() bool {
	return internal.IsProcessRunning(internal.FileshareBinaryPath)
}

type monitorFileshareProcessJob struct {
	isFileshareAllowed bool
	meshChecker        meshInternal.MeshnetChecker
	rulesController    rulesController
	processChecker     processChecker
}

type rulesController interface {
	ForbidFileshare() error
	PermitFileshare() error
}

type processChecker interface {
	isFileshareRunning() bool
}
