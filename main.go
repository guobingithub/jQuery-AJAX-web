package main

import (
	"jQuery-AJAX-web/logger"
	"fmt"
	"net/http"
	."jQuery-AJAX-web/controller"
)

func main() {
	logger.Info(fmt.Sprintf("jQuery-AJAX-web start..."))

	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/", AdminHandler)
	http.HandleFunc("/login/",LoginHandler)
	http.HandleFunc("/ajax/",AjaxHandler)
	http.HandleFunc("/",NotFoundHandler)
	http.ListenAndServe(":8808", nil)

}

