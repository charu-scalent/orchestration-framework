package service

import (
	"context"

	"github.com/scalent-io/orchestration-framework/entity"
	workflowContext "github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
	"github.com/scalent-io/orchestration-framework/pkg/log"
)

type WalletServiceImpl struct {
	walletRepo WalletRepo
}

func NewWalletServiceImpl(walletRepo WalletRepo) (*WalletServiceImpl, error) {
	return &WalletServiceImpl{walletRepo: walletRepo}, nil
}

func (s *WalletServiceImpl) Ping(ctx context.Context) errors.Response {
	reqID, _ := workflowContext.GetRequestIdFromContext(ctx)
	log.Info("wallet>wallet>service: pinf started ", reqID)

	s.walletRepo.Ping(ctx)
	return nil
}

func (s *WalletServiceImpl) GetWallet(ctx context.Context, token string) (*entity.User, errors.Response) {
	reqID, _ := workflowContext.GetRequestIdFromContext(ctx)
	log.Info("wallet>wallet>service: get wallet started", reqID)

	walletEntity, err := s.walletRepo.GetWallet(ctx, token)
	if err != nil {
		log.Error(err.Error(), reqID)
		return nil, err
	}

	log.Info("wallet>wallet>service: get wallet completed", reqID)
	return walletEntity, err
}
