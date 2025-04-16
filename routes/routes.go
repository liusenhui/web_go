package routes

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"runtime/debug"
	"web_go/pkg/logger"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(len(debug.Stack()) > 0))

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	// 注册业务路由组
	registerAPI(r)

	return r
}

// 获取http请求body的内容
func getBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()

	reader := bufio.NewReader(r.Body)
	chunk := make([]byte, 0, 200)
	buf := make([]byte, 200)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	return chunk, nil
}
