package router

import (
	"dahengzhang/news/db"
	"dahengzhang/news/dto"
	"dahengzhang/news/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
	preview := util.CutHTML(body.Content, 100)
	err := db.CreateNews(body.Title, preview, body.Content)
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
	preview := util.CutHTML(body.Content, 100)
	err = db.EditNews(id, body.Title, preview, body.Content)
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
