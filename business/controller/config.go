package controller

import (
	"net/http"

	"github.com/jeff-mendoza/integrator-dummy/business/model"
	"github.com/jeff-mendoza/integrator-dummy/business/repository"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context){
	request := model.CreateConfigRequest{}
	err := bindToModel(ctx, &request); if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,  gin.H{"message": "OK"})
	}

	err = repository.Set(request.CallerID, request.Token); if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,  gin.H{"message": err})
		return
	}
	response := model.CreateConfigResponse{CallerID: request.CallerID, Token: request.Token}
	ctx.IndentedJSON(http.StatusOK, response)
}

func Find(ctx *gin.Context){
	callerID := ctx.Param("caller-id")

	value, err := repository.Get(callerID); if err != nil {
		ctx.IndentedJSON(http.StatusNotFound,  gin.H{"message": "no-found"})
		return
	}

	response := model.CreateConfigResponse{CallerID: callerID, Token: value}

	ctx.IndentedJSON(http.StatusOK, response)
}




