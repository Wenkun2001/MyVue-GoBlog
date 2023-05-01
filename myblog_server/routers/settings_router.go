package routers

import (
	"myblog_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)
	router.PUT("settings", settingsApi.SettingsInfoUpdateView)
}
