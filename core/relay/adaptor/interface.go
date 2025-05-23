package adaptor

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/meta"
	relaymodel "github.com/labring/aiproxy/core/relay/model"
)

type Adaptor interface {
	GetBaseURL() string
	GetRequestURL(meta *meta.Meta) (string, error)
	SetupRequestHeader(meta *meta.Meta, c *gin.Context, req *http.Request) error
	ConvertRequest(meta *meta.Meta, req *http.Request) (method string, header http.Header, body io.Reader, err error)
	DoRequest(meta *meta.Meta, c *gin.Context, req *http.Request) (*http.Response, error)
	DoResponse(meta *meta.Meta, c *gin.Context, resp *http.Response) (*model.Usage, *relaymodel.ErrorWithStatusCode)
	GetModelList() []*model.ModelConfig
}

var ErrGetBalanceNotImplemented = errors.New("get balance not implemented")

type Balancer interface {
	GetBalance(channel *model.Channel) (float64, error)
}

type KeyValidator interface {
	ValidateKey(key string) error
	KeyHelp() string
}
