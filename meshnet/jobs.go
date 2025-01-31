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
	meshnetStatusPublisher events.PublishSubcriber[bool],
	meshnetEvents *Events,
) {
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

	// monitors the meshnet status and starts/stops the meshnet map refreshing job
	meshnetStatusPublisher.Subscribe(func(enabled bool) error {
		// TODO: check what happens if meshnet is started
		return startMeshnetRefresher(enabled, s.scheduler, s, s)
	})

	// monitor NC events
	meshnetEvents.PeerUpdate.Subscribe(func(s []string) error {
		
	})
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

func startMeshnetRefresher(enabled bool, scheduler gocron.Scheduler, meshnetChecker meshInternal.MeshnetChecker,
	fetcher meshInternal.MeshnetFetcher,
) error {
	if enabled {
		job, err := scheduler.NewJob(
			gocron.DurationJob(internal.MeshnetMapUpdateInterval),
			gocron.NewTask(jobs.JobRefreshMeshnetMap(meshnetChecker, fetcher)),
			gocron.WithName("refresh meshnet map"),
			gocron.WithTags(internal.MeshnetMapJobTag),
		)
		if err != nil {
			log.Println(internal.WarningPrefix, "job refresh meshnet schedule error:", err)
			return err
		}

		log.Println(internal.DebugPrefix, "meshnet map refresh job scheduled")

		if err := job.RunNow(); err != nil {
			log.Println(internal.ErrorPrefix, "failed to run meshnet map refresh job", err)
			return err
		}
	} else {
		scheduler.RemoveByTags(internal.MeshnetMapJobTag)
		log.Println(internal.DebugPrefix, "stop meshnet map refresh job")
	}
	return nil

}
