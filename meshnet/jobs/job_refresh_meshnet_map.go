package jobs

import (
	"log"

	"github.com/NordSecurity/nordvpn-linux/internal"
	meshInternal "github.com/NordSecurity/nordvpn-linux/meshnet/internal"
	"github.com/go-co-op/gocron/v2"
)

func ConfigureMeshnetMapRefresher(enabled bool, scheduler gocron.Scheduler, meshnetChecker meshInternal.MeshnetChecker,
	fetcher meshInternal.MeshnetFetcher,
) error {
	if enabled {
		job, err := scheduler.NewJob(
			gocron.DurationJob(internal.MeshnetMapUpdateInterval),
			gocron.NewTask(JobRefreshMeshnetMap(meshnetChecker, fetcher)),
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

func JobRefreshMeshnetMap(
	meshnetChecker meshInternal.MeshnetChecker,
	fetcher meshInternal.MeshnetFetcher,
) func() {
	return func() {
		if !meshnetChecker.IsMeshnetOn() {
			log.Println(internal.DebugPrefix, "updating meshnet map called when meshnet is not enabled")
			return
		}

		_, err := fetcher.RefreshMeshnetMap()
		if err != nil {
			log.Println(internal.ErrorPrefix, "job update meshnet map failed", err)
		}
	}
}
