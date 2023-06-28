package webapi

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Attach(r gin.IRouter)
}
