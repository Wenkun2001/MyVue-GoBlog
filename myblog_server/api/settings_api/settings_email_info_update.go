package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/config"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models/res"
)

func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
