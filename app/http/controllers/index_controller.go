package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Print("thie is controller init")
}

func Login(c *gin.Context) {
	fmt.Print("this is login")
}

func Register(c *gin.Context) {

}
