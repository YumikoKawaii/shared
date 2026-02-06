package interceptors

import (
	"context"
	"time"

	"github.com/YumikoKawaii/shared/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestID := uuid.New().String()
		startTime := time.Now()

		logger.Infof("[りくえすと] %s %s かいし", requestID, info.FullMethod)

		resp, err := handler(ctx, req)

		duration := time.Since(startTime)

		if err != nil {
			logger.Infof("[りくえすと] %s しっぱい (%v): %v", requestID, duration, err)
		} else {
			logger.Infof("[りくえすと] %s せいこう (%v)", requestID, duration)
		}

		return resp, err
	}
}

func StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		requestID := uuid.New().String()
		startTime := time.Now()

		logger.Infof("[すとりーむ] %s %s かいし", requestID, info.FullMethod)

		err := handler(srv, ss)

		duration := time.Since(startTime)

		if err != nil {
			logger.Infof("[すとりーむ] %s しっぱい (%v): %v", requestID, duration, err)
		} else {
			logger.Infof("[すとりーむ] %s せいこう (%v)", requestID, duration)
		}

		return err
	}
}
