package web

import (
	"fmt"
	"github.com/ratanraj/cryptochat"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewUserForm struct {
	Username  string `json:"username" form:"username"`
	PublicKey string `json:"publicKey" form:"publicKey"`
}

func RunServer(listenAddr string) error {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/newuser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newuser.html", gin.H{})
	})

	r.POST("/newuser", func(c *gin.Context) {
		var form NewUserForm
		err := c.Bind(&form)
		if err != nil {
			c.AbortWithError(403, fmt.Errorf("failed to authenticate"))
		}

		isKeyValid, err := cryptochat.VerifyPublicKey(form.PublicKey)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if isKeyValid {

		} else {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("invalid key"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"username": form.Username, "key": form.PublicKey})
	})

	return r.Run()
}
