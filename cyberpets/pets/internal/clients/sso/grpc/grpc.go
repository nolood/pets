package grpc

import (
	"context"
	"cyberpets/pets/internal/domain/sso"
	ssov1 "cyberpets/protos/gen/go/sso"
	"fmt"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	api ssov1.AuthClient
}

func New(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	const op = "clients.sso.grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{api: ssov1.NewAuthClient(cc)}, nil
}

func (c *Client) Validate(ctx context.Context, data sso.ValidateData) (bool, error) {
	const op = "clients.sso.grpc.Validate"

	resp, err := c.api.Validate(ctx, &ssov1.ValidateRequest{Token: data.Token, QueryId: data.QueryId, AuthDate: data.AuthDate, Hash: data.Hash, User: data.User})
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetOk(), nil
}

func (c *Client) ValidateToken(ctx context.Context, token string) (ok bool, userId string, err error) {

	const op = "clients.sso.grpc.ValidateToken"

	resp, err := c.api.ValidateToken(ctx, &ssov1.ValidateTokenRequest{Token: token})
	if err != nil {
		return false, "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetValid(), resp.GetUserId(), nil
}
