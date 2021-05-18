package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	secret = os.Getenv("SECRET")
	port = os.Getenv("PORT") 
)

type IntegratorChallengeResponse struct {
	EncryptedChallenge string `json:"encrypted_challenge"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal("pong")
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func webhookGetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("::webhook:validation\n")

	query, ok := r.URL.Query()["challengeCode"]
	if ok {
		challengeCode := query[0]
		fmt.Printf("\t::challengeCodeParam: " + challengeCode + "\n")
		sig := hmac.New(sha256.New, []byte(secret))
		sig.Write([]byte(challengeCode))
		model := IntegratorChallengeResponse{
			EncryptedChallenge: hex.EncodeToString(sig.Sum(nil)),
		}
		fmt.Printf("\t::response: " + model.EncryptedChallenge + "\n")
		fmt.Printf("=================================================================\n")
		jsonBytes, err := json.Marshal(model)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	} else {
		fmt.Printf("\t::fail\n")
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
	}

}

func webhookPostHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("::webhook:callback\n")
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Printf(string(bodyBytes) )
	fmt.Printf("\n\t::response: 200\n")
	fmt.Printf("=================================================================\n")
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func webhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		webhookGetHandler(w, r)
	case "POST":
		webhookPostHandler(w,r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}


func main() {
	if port == "" {
		port = ":8081"
        } else {
		port = ":" + port
	}

	fmt.Printf("::config\n")
	fmt.Printf("\t::secret: " + secret + "\n\n")
	fmt.Printf("\t::port: " + port + "\n\n")

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/webhook", webhookHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}
