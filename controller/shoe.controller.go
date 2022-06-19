package controller

import (
	"net/http"

	"example.com/sarang-apis/models"
	"example.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
)

type ShoeController struct {
	ShoeService services.ShoeService
}

func NewShoeServicew(shoeservice services.ShoeService) ShoeController {
	return ShoeController{
		ShoeService: shoeservice,
	}
}

func (uc *ShoeController) CreateShoe(ctx *gin.Context) {
	user_id := ctx.Param("id_user")
	var shoe models.Shoe

	if err := ctx.ShouldBindJSON(&shoe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.ShoeService.CreateShoe(&shoe, &user_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error() + user_id})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ShoeController) GetShoe(ctx *gin.Context) {
	shoebrand := ctx.Param("brand")
	user, err := uc.ShoeService.GetShoe(&shoebrand)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *ShoeController) GetAll(ctx *gin.Context) {
	shoes, err := uc.ShoeService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, shoes)
}

func (uc *ShoeController) UpdateShoe(ctx *gin.Context) {
	var shoe models.Shoe
	if err := ctx.ShouldBindJSON(&shoe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.ShoeService.UpdateShoe(&shoe)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ShoeController) DeleteShoe(ctx *gin.Context) {
	shoebrand := ctx.Param("brand")
	err := uc.ShoeService.DeleteShoe(&shoebrand)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		ctx.JSON(http.StatusBadGateway, ctx.Param("brand"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ShoeController) RegisterShoeRoutes(rg *gin.RouterGroup) {
	shorroute := rg.Group("/user")
	shorroute.POST("shoe/:id_user/create", uc.CreateShoe)
	shorroute.GET("shoe/get/:brand", uc.GetShoe)
	shorroute.GET("shoe/getall", uc.GetAll)
	shorroute.PATCH("shoe/:id_user/update", uc.UpdateShoe)
	shorroute.DELETE("shoe/:id_user/delete/:brand", uc.DeleteShoe)
}
