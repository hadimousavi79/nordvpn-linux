package internal

import "github.com/NordSecurity/nordvpn-linux/meshnet/pb"

type MeshnetChecker interface {
	IsMeshnetOn() bool
}

type MeshnetFetcher interface {
	FetchMeshnetPeers() *pb.GetPeersResponse
}

type MeshnetDataManager interface {
	GetMeshnetPeers() (*pb.GetPeersResponse, error)
	SetMeshnetPeers(peers *pb.GetPeersResponse, err error) bool
}
