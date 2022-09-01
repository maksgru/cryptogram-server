package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterUserHandlers(router *httprouter.Router) {
	router.POST("/users", createUserHandler)
}

type d struct {
	name string `json:"name"`
}

func createUserHandler(w http.ResponseWriter, req *http.Request, Params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(d{name: "max"})
	jd, _ := json.Marshal(d{name: "max"})
	w.Write(jd)
	fmt.Println(Params)
}
