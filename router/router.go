package router

import (
	"dahengzhang/news/db"
	"dahengzhang/news/dto"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var (
	tpl *template.Template
	err error
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index", nil)
}

func getNewsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := db.GetNews()
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  err.Error(),
		}
		rb.ReturnToFE(w)
		return
	}
	fmt.Print(res)
	rb := dto.ResBody{
		Code: 0,
		Data: res,
	}
	rb.ReturnToFE(w)
}

func createNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var body dto.News
	decoder.Decode(&body)
	err := db.CreateNews(body.Title, body.Content)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  err.Error(),
		}
		rb.ReturnToFE(w)
		return
	}

	rb := dto.ResBody{
		Code: 0,
		Msg:  "发布成功",
	}
	rb.ReturnToFE(w)
}

func editNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  "获取不到id",
		}
		rb.ReturnToFE(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var body dto.News
	decoder.Decode(&body)
	err = db.EditNews(id, body.Title, body.Content)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  err.Error(),
		}
		rb.ReturnToFE(w)
		return
	}
	rb := dto.ResBody{
		Code: 0,
		Msg:  "修改成功",
	}
	rb.ReturnToFE(w)
}

func searchNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  "获取不到id",
		}
		rb.ReturnToFE(w)
		return
	}
	res, err := db.SearchNews(id)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  err.Error(),
		}
		rb.ReturnToFE(w)
		return
	}
	rb := dto.ResBody{
		Code: 0,
		Data: res,
	}
	rb.ReturnToFE(w)
}

func deleteNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  "获取不到id",
		}
		rb.ReturnToFE(w)
		return
	}
	err = db.DeleteNews(id)
	if err != nil {
		rb := dto.ResBody{
			Code: 1,
			Msg:  err.Error(),
		}
		rb.ReturnToFE(w)
		return
	}
	rb := dto.ResBody{
		Code: 0,
		Data: "删除成功",
	}
	rb.ReturnToFE(w)
}

func init() {
	tpl, err = template.ParseGlob("tpl/*.html")
	if err != nil {
		log.Fatal(err.Error())
	}
}

// SetRouter 设置路由
func SetRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", home)
	router.GET("/news", getNewsList)
	router.GET("/news/:id", searchNews)
	router.POST("/news", createNews)
	router.PUT("/news/:id", editNews)
	router.DELETE("/news/:id", deleteNews)

	return router
}
