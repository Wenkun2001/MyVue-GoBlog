package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/config"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println("before", global.Config)
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
