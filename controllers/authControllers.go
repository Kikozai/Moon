package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/Kikozai/Moon/model"
	u "github.com/Kikozai/Moon/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	accound := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(accound)

	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
		return
	}
	resp := accound.Create()
	u.Respond(w, resp)
}


var Authenticate = func (w http.ResponseWriter, r *http.Request){ 
 
}