package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//res.Ok(map[string]string{}, "xxx", c)
	//c.JSON(200, gin.H{"msg": "xxx"})

	//res.OkWithData(global.Config.SiteInfo, c)
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "qi_niu":
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMessage("没有对应的配置参数信息", c)
	}
}
