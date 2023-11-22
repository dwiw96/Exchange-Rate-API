package api

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkContext(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Println("request is canceled")
		return status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline exceeded")
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	return nil
}
