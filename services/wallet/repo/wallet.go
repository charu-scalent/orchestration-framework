package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/scalent-io/orchestration-framework/entity"
	d2dContext "github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
	"github.com/scalent-io/orchestration-framework/pkg/log"
)

type WalletRepoImpl struct {
	db *sqlx.DB
}

func NewWalletRepoImpl(db *sqlx.DB) (*WalletRepoImpl, error) {
	return &WalletRepoImpl{db: db}, nil
}

func (r *WalletRepoImpl) Ping(ctx context.Context) errors.Response {
	reqID, _ := d2dContext.GetRequestIdFromContext(ctx)
	log.Info("Wallet repo: ping started", reqID)

	return nil
}

func (r *WalletRepoImpl) GetWallet(ctx context.Context, token string) (*entity.User, errors.Response) {
	reqID, _ := d2dContext.GetRequestIdFromContext(ctx)
	log.Info("wallet>wallet>repo: get wallet started ", reqID)
	return &entity.User{}, nil
}
