package user

import (
	. "go_server/handler"
	"go_server/model"
	"go_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}