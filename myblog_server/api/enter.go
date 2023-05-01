package api

import (
	"myblog_server/api/images_api"
	"myblog_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
