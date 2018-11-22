package controller

import (
	"net/http"
	"strings"
	"reflect"
	"html/template"
	"jQuery-AJAX-web/logger"
	"fmt"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("AdminHandler enter..."))

	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == ""{
		logger.Info("AdminHandler, no cookie, will Redirect to login.")
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	admin := &adminController{}
	controller := reflect.ValueOf(admin)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)

	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("AjaxHandler enter..."))

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	ajax := &ajaxController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)

	method.Call([]reflect.Value{responseValue, requestValue})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("LoginHandler enter..."))

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)

	method.Call([]reflect.Value{responseValue, requestValue})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("NotFoundHandler enter..."))

	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}

	t, err := template.ParseFiles("template/html/404.html")
	if err!=nil{
		logger.Error(fmt.Sprintf("NotFoundHandler, fail to template.ParseFiles, err:%v.",err))
	}

	t.Execute(w, nil)
}
