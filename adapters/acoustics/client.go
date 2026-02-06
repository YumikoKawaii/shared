package acoustics

import "context"

type Config struct {
	Endpoint string `json:"endpoint" mapstructure:"endpoint" yaml:"endpoint"`
	Protocol string `json:"protocol" mapstructure:"protocol" yaml:"protocol"`
}

func Initialize(config *Config) (Client, error) {
	switch config.Protocol {
	case "http":
		return newHTTPClient(config.Endpoint)
	default:
		return newGRPCClient(config.Endpoint)
	}
}

type EntryRequest struct{}

type EntryResponse struct{}

type Client interface {
	Entry(ctx context.Context, request EntryRequest) (EntryResponse, error)
}
