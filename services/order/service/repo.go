package service

import (
	"github.com/gin-gonic/gin"
	"github.com/scalent-io/orchestration-framework/entity"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
)

type OrderRepo interface {
	LoadOrder(ctx *gin.Context, userID, orderID string) (*entity.Order, errors.Response)
}
