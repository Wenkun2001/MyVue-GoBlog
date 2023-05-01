package routers

import "myblog_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
}
