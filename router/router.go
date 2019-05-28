package router

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	tpl *template.Template
	err error
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index", nil)
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

	router.POST("/user/register", registerUser)
	router.POST("/user/login", userLogin)
	router.GET("/user/check", checkJwt)

	return router
}
