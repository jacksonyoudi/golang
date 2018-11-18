package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)



/*
用户逻辑handler
 */



/*
  用户创建
 */
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "Create User Ha")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	userName := p.ByName("userName")
	io.WriteString(w, userName)
}