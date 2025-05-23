package render

import (
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/labring/aiproxy/core/common/conv"
)

func StringData(c *gin.Context, str string) {
	if len(c.Errors) > 0 {
		return
	}
	if c.IsAborted() {
		return
	}
	c.Render(-1, &OpenAISSE{Data: str})
	c.Writer.Flush()
}

func ObjectData(c *gin.Context, object any) error {
	if len(c.Errors) > 0 {
		return c.Errors.Last()
	}
	if c.IsAborted() {
		return errors.New("context aborted")
	}
	jsonData, err := sonic.Marshal(object)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}
	c.Render(-1, &OpenAISSE{Data: conv.BytesToString(jsonData)})
	c.Writer.Flush()
	return nil
}

const DONE = "[DONE]"

func Done(c *gin.Context) {
	StringData(c, DONE)
}
