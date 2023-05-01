package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/res"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
