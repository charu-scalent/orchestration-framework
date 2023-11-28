package service

import (
	"github.com/gin-gonic/gin"
	"github.com/scalent-io/orchestration-framework/entity"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
)

type OrderServiceImpl struct {
	orderRepo OrderRepo
}

func NewOrderServiceImpl(orderRepo OrderRepo) (*OrderServiceImpl, error) {
	if orderRepo == nil {
		return nil, errors.New("cartRepo dependency is nil")
	}

	return &OrderServiceImpl{orderRepo: orderRepo}, nil
}

func (s *OrderServiceImpl) Load(ctx *gin.Context, userID, orderID string) (*entity.Order, errors.Response) {
	orders, err := s.orderRepo.LoadOrder(ctx, userID, orderID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
