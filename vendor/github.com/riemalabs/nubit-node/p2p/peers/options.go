package peers

import (
	"fmt"
	"time"

	libhead "github.com/riemalabs/go-header"

	"github.com/riemalabs/nubit-node/strucs/eh"
	"github.com/riemalabs/nubit-node/p2p/p2psub"
)

type Parameters struct {
	// PoolValidationTimeout is the timeout used for validating incoming datahashes. Pools that have
	// been created for datahashes from p2psub that do not see this hash from headersub after this
	// timeout will be garbage collected.
	PoolValidationTimeout time.Duration

	// PeerCooldown is the time a peer is put on cooldown after a ResultCooldownPeer.
	PeerCooldown time.Duration

	// GcInterval is the interval at which the manager will garbage collect unvalidated pools.
	GcInterval time.Duration

	// EnableBlackListing turns on blacklisting for misbehaved peers
	EnableBlackListing bool
}

type Option func(*Manager) error

// Validate validates the values in Parameters
func (p *Parameters) Validate() error {
	if p.PoolValidationTimeout <= 0 {
		return fmt.Errorf("peer-manager: validation timeout must be positive")
	}

	if p.PeerCooldown <= 0 {
		return fmt.Errorf("peer-manager: peer cooldown must be positive")
	}

	if p.GcInterval <= 0 {
		return fmt.Errorf("peer-manager: garbage collection interval must be positive")
	}

	return nil
}

// DefaultParameters returns the default configuration values for the peer manager parameters
func DefaultParameters() Parameters {
	return Parameters{
		// PoolValidationTimeout's default value is based on the default daser sampling timeout of 1 minute.
		// If a received datahash has not tried to be sampled within these two minutes, the pool will be
		// removed.
		PoolValidationTimeout: 2 * time.Minute,
		// PeerCooldown's default value is based on initial network tests that showed a ~3.5 second
		// sync time for large blocks. This value gives our (discovery) peers enough time to sync
		// the new block before we ask them again.
		PeerCooldown: 3 * time.Second,
		GcInterval:   time.Second * 30,
		// blacklisting is off by default //TODO(@walldiss): enable blacklisting once all related issues
		// are resolved
		EnableBlackListing: false,
	}
}

// WithP2pSubPools passes a p2psub and headersub instance to be used to populate and validate
// pools from p2psub notifications.
func WithP2pSubPools(p2pSub *p2psub.PubSub, headerSub libhead.Subscriber[*header.ExtendedHeader]) Option {
	return func(m *Manager) error {
		m.p2pSub = p2pSub
		m.headerSub = headerSub
		return nil
	}
}

// WithMetrics turns on metric collection in peer manager.
func (m *Manager) WithMetrics() error {
	metrics, err := initMetrics(m)
	if err != nil {
		return fmt.Errorf("peer-manager: init metrics: %w", err)
	}
	m.metrics = metrics
	return nil
}
