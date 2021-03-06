package v1

import (
	"github.com/lairdnote/social/model"

	"github.com/gin-gonic/gin"
)

// 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if userID, ok := c.Get("user_id"); ok {
		if user, err := model.GetUser(*userID.(*uint)); err == nil {
			return user
		}
	}
	return nil
}
