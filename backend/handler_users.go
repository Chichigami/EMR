package main

import (
	"github.com/gin-gonic/gin"
)

func (cfg *Config) handlerUsersNew(c *gin.Context) {
	//make new login, admin perm
	handlerPlaceholder(c)
}

func (cfg *Config) handlerUsersVerify(c *gin.Context) {
	//from POST req
	// username := "get username"
	// password := "get password from req"
	// err := cfg.db.GetHashedPassword(c.Request.Context(), username)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, gin.H{
	// 		"success": false,
	// 		"message": err.Error(),
	// 	})
	// }
	// if success := auth.CheckPasswordHash(cfg.db.GetHashedPassword(username), password); success == nil {
	// 	//give auth cookie
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"success": true,
	// 		"message": "Login successful",
	// 	})
	// }

	handlerPlaceholder(c)
}

func (cfg *Config) handlerUsersUpdate(c *gin.Context) {
	//after verified user, can update user or password. if admin can update other perms.
	handlerPlaceholder(c)
}

func (cfg *Config) handlerUsersDelete(c *gin.Context) {
	//delete login, admin perm
	handlerPlaceholder(c)
}
