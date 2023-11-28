package repo

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/scalent-io/orchestration-framework/entity"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
)

type OrderRepoImpl struct {
	db *sqlx.DB
}

func NewOrderRepoImpl(db *sqlx.DB) (*OrderRepoImpl, error) {
	return &OrderRepoImpl{
		db: db,
	}, nil
}

func (r *OrderRepoImpl) LoadOrder(ctx *gin.Context, userID, orderID string) (*entity.Order, errors.Response) {

	return &entity.DummyOrders[0], nil
}
