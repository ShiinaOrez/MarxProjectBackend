package cpc

import (
	"github.com/ShiinaOrez/MarxProjectBackend/handler"
	"github.com/ShiinaOrez/MarxProjectBackend/model"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/constvar"
	"github.com/ShiinaOrez/MarxProjectBackend/pkg/errno"
	"github.com/gin-gonic/gin"
)

type CpcArticlesResponse struct {
	Articles  []model.ArticlesModel `json:"articles"`
	PageCount int                   `json:"page_count"`
}

func StudyCpcArticles(c *gin.Context) {
	page := c.GetInt("page")
	limit := constvar.DefaultLimit
	articles, count, err := model.GetArticles(page, limit)
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}
	pageCount := int(count) / limit
	if int(count)%limit != 0 {
		pageCount += 1
	}
	response := CpcArticlesResponse{
		Articles:  articles,
		PageCount: pageCount,
	}
	handler.SendResponse(c, nil, response)
	return
}

func StudyCpcArticle(c *gin.Context) {
	id := c.GetInt("ID")
	article, err := model.GetArticleById(id)
	if err != nil {
		handler.SendError(c, errno.ErrArticleNotFound, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *article)
}
