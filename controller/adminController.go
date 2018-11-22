package controller

import (
    "net/http"
    "html/template"
	"jQuery-AJAX-web/logger"
	"fmt"
)

type User struct {
    UserName string
}

type adminController struct {

}

func (this *adminController)IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	logger.Error(fmt.Sprintf("adminController IndexAction enter, r:%v.",r))

    t, err := template.ParseFiles("template/html/admin/index.html")
    if err!=nil{
        logger.Error(fmt.Sprintf("adminController IndexAction, fail to template.ParseFiles, err:%v.",err))
    }

    t.Execute(w, &User{user})
}