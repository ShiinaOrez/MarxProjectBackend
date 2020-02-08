package study

import (
	"github.com/ShiinaOrez/MarxProjectBackend/handler"
	"github.com/ShiinaOrez/MarxProjectBackend/model"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/constvar"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/errno"
	"github.com/gin-gonic/gin"
)

type StudyNewsResponse struct {
	News      []model.NewsModel `json:"news"`
	PageCount int               `json:"page_count"`
}

func StudyNews(c *gin.Context) {
	page := c.GetInt("page")
	limit := constvar.DefaultLimit
	news, count, err := model.GetStudyNews(page, limit)
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}
	pageCount := int(count) / limit
	if int(count)%limit != 0 {
		pageCount += 1
	}
	response := StudyNewsResponse{
		News:      news,
		PageCount: pageCount,
	}
	handler.SendResponse(c, nil, response)
	return
}

func StudyNew(c *gin.Context) {
	id := c.GetInt("ID")
	new, err := model.GetStudyNewById(id)
	if err != nil {
		handler.SendError(c, errno.ErrNewNotFound, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *new)
}
