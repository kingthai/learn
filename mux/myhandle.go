package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MyGetHandler(w http.ResponseWriter, r *http.Request) {
	// parse query parameter
	//vals := r.URL.Query()
	//param, _ := vals["servicename"]  // get query parameters

	params := r.URL.Query()
	imageName := params.Get("image")
	namespace := params.Get("namespace")
	secretName := params.Get("secret")
	fmt.Println(imageName, namespace, secretName)

	// composite response body
	//var res = map[string]string{"result":"succ", "name":param[0]}
	//response, _ := json.Marshal(res)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(response)
}

func MyPostHandler(w http.ResponseWriter, r *http.Request) {
	// parse path variable


	//servicename := vars["servicename"]

	// form data
	a := r.FormValue("servicename")
	b := r.FormValue("bbbb")
	fmt.Println(a, b)

	// parse JSON body
	//var req map[string]interface{}
	//body, _ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(body, &req)
	//servicetype := req["servicetype"].(string)

	// composite response body
	//var res = map[string]string{"result":"succ", "name":"1", "type":servicetype}
	res:=map[string]string{"result":"123"}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
