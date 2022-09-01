package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// get list of tags
		apiv1.GET("/tags", v1.GetTags)
		// new tag
		apiv1.POST("/tags", v1.AddTag)
		// edit tag
		apiv1.PUT("/tags/:id", v1.EditTag)
		// delete tag
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// export tag
		r.POST("/tags/export", v1.ExportTag)
		// import tag
		r.POST("/tags/import", v1.ImportTag)

		// get list of articles
		apiv1.GET("/articles", v1.GetArticles)
		// get article by id
		apiv1.GET("/articles/:id", v1.GetArticle)
		// add article
		apiv1.POST("/articles", v1.AddArticle)
		// edit article
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//delete article
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		// generate article poster
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
