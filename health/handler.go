package health

import (
	"context"
	"time"

	api "github.com/YumikoKawaii/rpc.com/protobuf/health"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	api.UnimplementedHealthServer
}

func Initialize() *Handler {
	logger.Info("[へるす] はんどらーしょきかかんりょう")
	return &Handler{}
}

func (h *Handler) Liveness(ctx context.Context, request *api.LivenessRequest) (*api.LivenessResponse, error) {
	return &api.LivenessResponse{
		Alive:     true,
		Message:   "ゆみこせかいでいちばんかわいい",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (h *Handler) Readiness(ctx context.Context, request *api.ReadinessRequest) (*api.ReadinessResponse, error) {
	return &api.ReadinessResponse{
		Ready:     true,
		Message:   "あとりあばんざいいい",
		Timestamp: time.Now().Unix(),
	}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterHealthServer(instance.Instance(), h)
	return api.RegisterHealthHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
}
