package history

import (
	"github.com/ShiinaOrez/MarxProjectBackend/handler"
	"github.com/ShiinaOrez/MarxProjectBackend/model"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/constvar"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/errno"
	"github.com/gin-gonic/gin"
)

type HistoryNewsResponse struct {
	News      []model.NewsModel `json:"news"`
	PageCount int               `json:"page_count"`
}

func HistoryNews(c *gin.Context) {
	page := c.GetInt("page")
	limit := constvar.DefaultLimit
	news, count, err := model.GetHistoryNews(page, limit)
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}
	pageCount := int(count) / limit
	if int(count)%limit != 0 {
		pageCount += 1
	}
	response := HistoryNewsResponse{
		News:      news,
		PageCount: pageCount,
	}
	handler.SendResponse(c, nil, response)
	return
}

func HistoryNew(c *gin.Context) {
	id := c.GetInt("ID")
	new, err := model.GetHistoryNewById(id)
	if err != nil {
		handler.SendError(c, errno.ErrNewNotFound, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *new)
}
