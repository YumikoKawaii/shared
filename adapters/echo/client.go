package acoustics

import "context"

type Config struct {
	Endpoint string `json:"endpoint" mapstructure:"endpoint" yaml:"endpoint"`
	Protocol string `json:"protocol" mapstructure:"protocol" yaml:"protocol"`
}

type ChargeRequest struct{}

type ChargeResponse struct{}

type DischargeRequest struct{}

type DischargeResponse struct{}

type Client interface {
	Charge(ctx context.Context, request ChargeRequest) (ChargeResponse, error)
	Discharge(ctx context.Context, request DischargeRequest) (DischargeResponse, error)
}
