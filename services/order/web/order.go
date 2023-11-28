package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apimodel "github.com/scalent-io/orchestration-framework/apimodel/order"
	"github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/log"
	httpUtils "github.com/scalent-io/orchestration-framework/pkg/utils"
)

func (s OrderHandlerRegistry) GetOrderHandler(c *gin.Context) {
	reqID, _ := context.GetRequestIdFromContext(c.Request.Context())
	log.Info("user>order>web: get order started", reqID)

	var req apimodel.GetOrderRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.ErrorWithData("Insufficient Parameters --> ", err.Error(), reqID)
		return
	}

	orderEntity, err := s.options.OrderService.Load(c, req.UserID, req.OrderID)
	if err != nil {
		log.Error(err.Error(), reqID)
		return
	}

	log.Info("user>order>web: get order completed", reqID)
	httpUtils.DataResponse(c, http.StatusOK, "order details fetched successfully", orderEntity)
}
