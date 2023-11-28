package service

import (
	"context"

	"github.com/scalent-io/orchestration-framework/entity"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
)

type WalletService interface {
	Ping(ctx context.Context) errors.Response

	GetWallet(ctx context.Context, token string) (*entity.User, errors.Response)
}
