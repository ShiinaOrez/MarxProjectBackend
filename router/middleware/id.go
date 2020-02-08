package middleware

import (
	"github.com/ShiinaOrez/MarxProjectBackend/handler"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Id(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrIdInvalid, nil, err.Error())
		c.AbortWithStatus(400)
	}
	c.Set("ID", ID)
	c.Next()
}
