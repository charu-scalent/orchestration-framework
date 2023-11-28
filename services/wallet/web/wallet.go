package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/log"
	httpUtils "github.com/scalent-io/orchestration-framework/pkg/utils"
)

func (s UserHandlerRegistry) homeHandler(c *gin.Context) {
	reqID, _ := context.GetRequestIdFromContext(c.Request.Context())
	log.Info("user>user>web: ping started", reqID)

	s.options.WalletService.Ping(c)

	log.Info("user>user>web: ping completed", reqID)
	httpUtils.DataResponse(c, http.StatusOK, "Ping successful", nil)

}

func (s UserHandlerRegistry) GetWalletHandler(c *gin.Context) {
	reqID, _ := context.GetRequestIdFromContext(c.Request.Context())
	log.Info("user>wallet>web: get wallet started", reqID)

	token := "token1"

	walletEntity, err := s.options.WalletService.GetWallet(c, token)
	if err != nil {
		log.Error(err.Error(), reqID)
		return
	}

	log.Info("user>wallet>web: get wallet completed", reqID)
	httpUtils.DataResponse(c, http.StatusOK, "Wallet details fetched successfully", walletEntity)

}
