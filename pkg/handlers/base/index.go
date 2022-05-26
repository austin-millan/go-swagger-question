package base

import (
	"github.com/austin-millan/go-swagger-question/pkg/server/base/models"
	"github.com/austin-millan/go-swagger-question/pkg/server/base/operations/index"
	"github.com/go-openapi/runtime/middleware"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (handler PingHandler) Handle(params index.IndexParams) middleware.Responder {
	return &index.IndexOK{
		Payload: &models.PingResponse{
			ResponseMeta: models.ResponseMeta{
				Meta: "meta",
			},
		},
	}
}
