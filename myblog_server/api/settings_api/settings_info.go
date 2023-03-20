package settings_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/models/res"
)

func (SettingsApi) SetttingsInfoInfoView(c *gin.Context) {
	res.Ok(map[string]string{}, "xxx", c)
	//c.JSON(200, gin.H{"msg": "xxx"})
}
