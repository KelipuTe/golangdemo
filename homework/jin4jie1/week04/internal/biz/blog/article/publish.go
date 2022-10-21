package article

import (
	pkg_context "demo-golang/homework/jin4jie1/week04/internal/pkg/context"
	blog_article "demo-golang/homework/jin4jie1/week04/internal/service/blog/article"
	"demo_golang/go_homework/week04/internal/biz/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PublishArticle(p1gc *gin.Context) {
	// 一通操作从鉴权是中间件的处理结果中获取用户id
	userId := int64(1)
	// 一通操作把article的内容搞出来
	p1bm := &blog_article.ArticleBizModel{
		Content: "article content",
	}
	bc := pkg_context.MakeBlogContext(p1gc)
	articleId, err := blog_article.PublishArticle(bc, userId, p1bm)
	if nil != err {
		response.Set500Response(p1gc, fmt.Sprintf("%s", err))
		return
	}
	response.Set200Response(p1gc, map[string]interface{}{"article": articleId})
}
