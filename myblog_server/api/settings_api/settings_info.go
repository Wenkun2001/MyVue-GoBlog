package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//res.Ok(map[string]string{}, "xxx", c)
	//c.JSON(200, gin.H{"msg": "xxx"})
	res.OkWithData(global.Config.SiteInfo, c)
}
