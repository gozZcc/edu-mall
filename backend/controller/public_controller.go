package controller

import "github.com/gin-gonic/gin"

type PublicController struct {
}

func NewPublicController() *PublicController {
	return &PublicController{}
}

func (p *PublicController) Hello(ctx *gin.Context) {

}
