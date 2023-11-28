package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/scalent-io/orchestration-framework/entity"
	workflowContext "github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
	"github.com/scalent-io/orchestration-framework/pkg/log"
)

var (
	Users = []entity.User{
		{ID: 1, Name: "Charu", WalletBalance: 1000, Token: "token1"},
		{ID: 2, Name: "Swati", WalletBalance: 1500, Token: "token2"},
		{ID: 3, Name: "Prajkta", WalletBalance: 800, Token: "token3"},
		{ID: 4, Name: "Akash", WalletBalance: 1200, Token: "token4"},
	}
)

type WalletRepoImpl struct {
	db *sqlx.DB
}

func NewWalletRepoImpl(db *sqlx.DB) (*WalletRepoImpl, error) {
	return &WalletRepoImpl{db: db}, nil
}

func (r *WalletRepoImpl) Ping(ctx context.Context) errors.Response {
	reqID, _ := workflowContext.GetRequestIdFromContext(ctx)
	log.Info("Wallet repo: ping started", reqID)

	return nil
}

func (r *WalletRepoImpl) GetWallet(ctx context.Context, token string) (*entity.User, errors.Response) {
	reqID, _ := workflowContext.GetRequestIdFromContext(ctx)
	log.Info("wallet>wallet>repo: get wallet started ", reqID)

	var result entity.User
	for _, user := range Users {
		if user.Token == "token1" {
			result = user
			break
		}
	}
	return &result, nil
}
