package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeff-mendoza/integrator-dummy/business/model"
	"github.com/jeff-mendoza/integrator-dummy/business/repository"

	"github.com/gin-gonic/gin"
)

type IntegratorChallengeResponse struct {
	EncryptedChallenge string `json:"encrypted_challenge"`
}


func Check(ctx *gin.Context){
	fmt.Printf("::webhook:validation")

	callerID := ctx.Param("caller-id")
	challengeCode, ok := ctx.GetQuery("challengeCode")

	if ok {
		token, err := repository.Get(callerID); if err != nil {
			ctx.IndentedJSON(http.StatusNotFound,  gin.H{"message": "no-found"})
			return
		}
		sig := hmac.New(sha256.New, []byte(token))
		sig.Write([]byte(challengeCode))
		response := IntegratorChallengeResponse{
			EncryptedChallenge: hex.EncodeToString(sig.Sum(nil)),
		}
		ctx.IndentedJSON(http.StatusOK, response)
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error"})
	}
}

func Callback(ctx *gin.Context){
	fmt.Printf("::webhook:callback")
	request := model.WebhookRequest{}
	err := bindToModel(ctx, &request); if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,  gin.H{"message": "OK"})
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,  gin.H{"message": err})
		return
	}

	err = repository.Set(request.ID, string(requestBytes)); if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,  gin.H{"message": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK,  gin.H{"message": "OK"})
}

func FindEvent(ctx *gin.Context){
	fmt.Printf("::webhook:find-event")

	paymentIntentID := ctx.Param("payment-intent-id")
	event, err := repository.Get(paymentIntentID); if err != nil {
		ctx.IndentedJSON(http.StatusNotFound,  gin.H{"message": "no-found"})
		return
	}

	response := model.WebhookRequest{}
	err = json.Unmarshal([]byte(event), &response); if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,  gin.H{"message": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func Ping(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK,  gin.H{"message": "pong"})
}


func bindToModel(context *gin.Context, model interface{}) error {
	if err := context.ShouldBindJSON(model); err != nil {
		return fmt.Errorf("bind error")
	}
	return nil
}
