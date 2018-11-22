package controller

import (
    "net/http"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/thrsafe"
    "encoding/json"
	"jQuery-AJAX-web/logger"
	"fmt"
)

const (
	db_addr = "127.0.0.1:3306"
	db_name = "webdemo"
	db_user = "root"
	db_pass = "123456"
)

type Result struct{
    Ret int
    Reason string
    Data interface{}
}

type ajaxController struct {

}

func (this *ajaxController)LoginAction(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("ajaxController LoginAction enter, r:%v.",r))

    w.Header().Set("content-type", "application/json")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "参数错误", nil)
        return
    }
    
    admin_name := r.FormValue("admin_name")
    admin_password := r.FormValue("admin_password")
    
    if admin_name == "" || admin_password == ""{
        OutputJson(w, 0, "参数错误", nil)
        return
    }
    
    db := mysql.New("tcp", "", db_addr, db_user, db_pass, db_name)
    if err := db.Connect(); err != nil {
        logger.Error(fmt.Sprintf("ajaxController LoginAction, fail to Connect mysql, err:%v.",err))
        OutputJson(w, 0, "数据库操作失败", nil)
        return
    }
    defer db.Close()
    
    rows, res, err := db.Query("select * from webdemo_admin where admin_name = '%s'", admin_name)
    if err != nil {
        logger.Error(fmt.Sprintf("ajaxController LoginAction, fail to Query, err:%v.",err))
        OutputJson(w, 0, "数据库操作失败", nil)
        return
    }
    
    name := res.Map("admin_password")
    admin_password_db := rows[0].Str(name)
    
    if admin_password_db != admin_password {
        OutputJson(w, 0, "密码输入错误", nil)
        return
    }
    
    // 存入cookie,使用cookie存储
    cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
    http.SetCookie(w, &cookie)
    
    OutputJson(w, 1, "操作成功", nil)
    return
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}