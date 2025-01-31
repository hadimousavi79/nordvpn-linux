package jobs

import (
	"log"

	"github.com/NordSecurity/nordvpn-linux/events"
	"github.com/NordSecurity/nordvpn-linux/internal"
	meshInternal "github.com/NordSecurity/nordvpn-linux/meshnet/internal"
)

func JobRefreshMeshnetMap(
	meshnetChecker meshInternal.MeshnetChecker,
	fetcher meshInternal.MeshnetFetcher,
	dm meshInternal.MeshnetDataManager,
	publisher events.Publisher[[]string],
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
