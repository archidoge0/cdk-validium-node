package dataavailability

// DABackendType is the data availability protocol for the CDK
type DABackendType string

const (
	// DataAvailabilityCommittee is the DAC protocol backend
	DataAvailabilityCommittee DABackendType = "DataAvailabilityCommittee"
	// DataAvailabilityNubitDA is the NubitDA protocol backend
	DataAvailabilityNubitDA DABackendType = "Nubit"
)

type Config struct {
	NodeRPC   string `mapstructure:"NodeRPC"`
	AuthToken string `mapstructure:"AuthToken"`
	Namespace string `mapstructure:"Namespace"`
}
