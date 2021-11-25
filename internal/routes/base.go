package routes

import (
	"github.com/gin-gonic/gin"
)

type RouterInterface interface {
	Initialize(engine *gin.Engine)
}

// 注册路由
type Routers struct {
	*Tag
	*Artists
}

func (r *Routers) SetupRouter(engine *gin.Engine) *gin.Engine {
	//注册用户路由
	r.Tag.Initialize(engine)
	r.Artists.Initialize(engine)
	return engine
}

func NewRouters() *Routers {
	return &Routers{
		Tag: &Tag{},
		Artists: &Artists{},
	}
}