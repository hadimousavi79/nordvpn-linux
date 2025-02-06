package mock

import (
	"github.com/NordSecurity/nordvpn-linux/config"
	daemonevents "github.com/NordSecurity/nordvpn-linux/daemon/events"
	"github.com/NordSecurity/nordvpn-linux/events"
)

func NewSettingsEmptyEvents() *daemonevents.SettingsEvents {
	return &daemonevents.SettingsEvents{
		Killswitch:           &daemonevents.MockPublisherSubscriber[bool]{},
		Autoconnect:          &daemonevents.MockPublisherSubscriber[bool]{},
		DNS:                  &daemonevents.MockPublisherSubscriber[events.DataDNS]{},
		ThreatProtectionLite: &daemonevents.MockPublisherSubscriber[bool]{},
		Protocol:             &daemonevents.MockPublisherSubscriber[config.Protocol]{},
		Allowlist:            &daemonevents.MockPublisherSubscriber[events.DataAllowlist]{},
		Technology:           &daemonevents.MockPublisherSubscriber[config.Technology]{},
		Obfuscate:            &daemonevents.MockPublisherSubscriber[bool]{},
		Firewall:             &daemonevents.MockPublisherSubscriber[bool]{},
		Routing:              &daemonevents.MockPublisherSubscriber[bool]{},
		Notify:               &daemonevents.MockPublisherSubscriber[bool]{},
		Meshnet:              &daemonevents.MockPublisherSubscriber[bool]{},
		Ipv6:                 &daemonevents.MockPublisherSubscriber[bool]{},
		Defaults:             &daemonevents.MockPublisherSubscriber[any]{},
		LANDiscovery:         &daemonevents.MockPublisherSubscriber[bool]{},
		VirtualLocation:      &daemonevents.MockPublisherSubscriber[bool]{},
		PostquantumVPN:       &daemonevents.MockPublisherSubscriber[bool]{},
	}
}
