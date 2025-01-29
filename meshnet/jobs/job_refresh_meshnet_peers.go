package jobs

import (
	"log"

	"github.com/NordSecurity/nordvpn-linux/events"
	"github.com/NordSecurity/nordvpn-linux/internal"
	meshInternal "github.com/NordSecurity/nordvpn-linux/meshnet/internal"
)

func JobUpdateMeshnetPeers(
	meshnetChecker meshInternal.MeshnetChecker,
	fetcher meshInternal.MeshnetFetcher,
	dm meshInternal.MeshnetDataManager,
	publisher events.Publisher[[]string],
) func() {
	return func() {
		if !meshnetChecker.IsMeshnetOn() {
			log.Println(internal.DebugPrefix, "updating meshnet peers when it is not enabled")
			return
		}

		response := fetcher.FetchMeshnetPeers()
		if dm.SetMeshnetPeers(response, nil) {
			// notify that peers changed
			publisher.Publish(nil)
			log.Println(internal.DebugPrefix, "notify meshnet peers changed")
		}
	}
}

// func fetchMeshnetPeers(
// 	ac auth.Checker,
// 	cm config.Manager,
// 	mc meshnet.Checker,
// 	reg mesh.Registry,
// 	pub events.Publisher[error],
// 	meshnetChecker meshnet.MeshnetChecker,
// 	netw meshnet.Networker,
// 	daemonEvents *daemonevents.Events,
// ) (*pb.PeerList, error) {
// 	if !ac.IsLoggedIn() {
// 		return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{ServiceErrorCode: pb.ServiceErrorCode_NOT_LOGGED_IN})
// 	}

// 	var cfg config.Config
// 	if err := cm.Load(&cfg); err != nil {
// 		pub.Publish(err)
// 		return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 			ServiceErrorCode: pb.ServiceErrorCode_CONFIG_FAILURE,
// 		})
// 	}

// 	if !cfg.Mesh {
// 		return nil, NewGenericError(&pb.GetPeersResponse_MeshnetErrorCode{
// 			MeshnetErrorCode: pb.MeshnetErrorCode_NOT_ENABLED,
// 		})
// 	}

// 	peers := pb.PeerList{}

// 	if !mc.IsRegistrationInfoCorrect() {
// 		token := cfg.TokensData[cfg.AutoConnectData.ID].Token
// 		resp, err := reg.Local(token)
// 		if err != nil {
// 			if errors.Is(err, core.ErrUnauthorized) {
// 				if err := cm.SaveWith(auth.Logout(cfg.AutoConnectData.ID, daemonEvents.User.Logout)); err != nil {
// 					pub.Publish(err)
// 					return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 						ServiceErrorCode: pb.ServiceErrorCode_CONFIG_FAILURE,
// 					})
// 				}
// 				return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 					ServiceErrorCode: pb.ServiceErrorCode_NOT_LOGGED_IN,
// 				})
// 			}
// 			pub.Publish(fmt.Errorf("listing local peers (@GetPeers): %w", err))

// 			// Mesh could get disabled (when self is removed)
// 			//  - check it and report it to the user properly.
// 			if !meshnetChecker.IsMeshnetOn() {
// 				return nil, NewGenericError(&pb.GetPeersResponse_MeshnetErrorCode{
// 					MeshnetErrorCode: pb.MeshnetErrorCode_NOT_ENABLED,
// 				})
// 			}

// 			return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 				ServiceErrorCode: pb.ServiceErrorCode_API_FAILURE,
// 			})
// 		}

// 		for _, peer := range resp {
// 			peers.Local = append(peers.Local, peer.ToProtobuf())
// 		}
// 	} else {
// 		token := cfg.TokensData[cfg.AutoConnectData.ID].Token
// 		resp, err := reg.List(token, cfg.MeshDevice.ID)
// 		if err != nil {
// 			if errors.Is(err, core.ErrUnauthorized) {
// 				if err := cm.SaveWith(auth.Logout(cfg.AutoConnectData.ID, daemonEvents.User.Logout)); err != nil {
// 					pub.Publish(err)
// 					return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 						ServiceErrorCode: pb.ServiceErrorCode_CONFIG_FAILURE,
// 					})
// 				}
// 				return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 					ServiceErrorCode: pb.ServiceErrorCode_NOT_LOGGED_IN,
// 				})
// 			}
// 			pub.Publish(fmt.Errorf("listing peers (@GetPeers): %w", err))

// 			// Mesh could get disabled (when self is removed)
// 			//  - check it and report it to the user properly.
// 			if !meshnetChecker.IsMeshnetOn() {
// 				return nil, NewGenericError(&pb.GetPeersResponse_MeshnetErrorCode{
// 					MeshnetErrorCode: pb.MeshnetErrorCode_NOT_ENABLED,
// 				})
// 			}

// 			return nil, NewGenericError(&pb.GetPeersResponse_ServiceErrorCode{
// 				ServiceErrorCode: pb.ServiceErrorCode_API_FAILURE,
// 			})
// 		}

// 		peers.Self = cfg.MeshDevice.ToProtobuf()
// 		peerMap, err := netw.StatusMap()
// 		if err != nil {
// 			peerMap = map[string]string{}
// 		}
// 		for _, peer := range resp {
// 			protoPeer := peer.ToProtobuf()
// 			status := pb.PeerStatus_DISCONNECTED
// 			if peerMap[peer.PublicKey] == "connected" {
// 				status = pb.PeerStatus_CONNECTED
// 			}
// 			protoPeer.Status = status
// 			if peer.IsLocal {
// 				peers.Local = append(peers.Local, protoPeer)
// 			} else {
// 				peers.External = append(peers.External, protoPeer)
// 			}
// 		}
// 	}

// 	if s.lastPeers != peers.String() {
// 		s.lastPeers = peers.String()
// 		s.subjectPeerUpdate.Publish(nil)
// 	}

// 	return &peers, nil
// }
