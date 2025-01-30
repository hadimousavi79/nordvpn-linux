package internal

import (
	"github.com/NordSecurity/nordvpn-linux/core/mesh"
	"github.com/NordSecurity/nordvpn-linux/meshnet/pb"
)

type MeshnetChecker interface {
	IsMeshnetOn() bool
}

type MeshnetFetcher interface {
	RefreshMeshnetMap() *pb.GetPeersResponse
}

type MeshnetDataManager interface {
	GetMeshnetMap() (*mesh.MachineMap, error)
	SetMeshnetMap(peers *mesh.MachineMap, err error) bool
}
