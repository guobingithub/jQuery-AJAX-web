package controller

import (
    "net/http"
    "html/template"
	"jQuery-AJAX-web/logger"
	"fmt"
)

type loginController struct {

}

func (this *loginController)IndexAction(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("loginController IndexAction enter, r:%v.",r))

    t, err := template.ParseFiles("template/html/login/index.html")
    if err!=nil{
        logger.Error(fmt.Sprintf("loginController IndexAction, fail to template.ParseFiles, err:%v.",err))
    }

    t.Execute(w, nil)
}