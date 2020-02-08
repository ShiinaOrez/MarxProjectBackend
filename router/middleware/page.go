package middleware

import (
	"github.com/ShiinaOrez/MarxProjectBackend/handler"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Page(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrPageInvalid, nil, err.Error())
		c.AbortWithStatus(400)
	}
	c.Set("page", page)
	c.Next()
}
