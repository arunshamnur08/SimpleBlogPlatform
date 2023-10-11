package routers

import (
	"github.com/gin-gonic/gin"
	"simple-blogging-platform/internal"
)

func ApiRouters(r *gin.Engine, blog_client internal.Blog) {
	r.POST("blog/Create", blog_client.Createblog)
	r.GET("blog/GetAllBlogs", blog_client.GetAllBlogs)
	r.GET("blog/:blogId", blog_client.GetBlogById)
	r.PATCH("blog/:blogId", blog_client.UpdateBlogById)
	r.DELETE("blog/:blogId", blog_client.DeleteBlogBYID)
}
