package router

import (
	"dahengzhang/news/dto"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func registerUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rb := dto.ResBody{
		Code: 0,
		Msg:  "注册成功",
	}
	rb.ReturnToFE(w)
}

func userLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Date(2018, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("hmacSampleSecret"))
	if err != nil {
		fmt.Print("token error", err.Error())
	}
	rb := dto.ResBody{
		Code: 0,
		Msg:  "登陆成功",
		Data: tokenString,
	}
	rb.ReturnToFE(w)
}

func checkJwt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) { // sample token string taken from the New example
	vars := r.URL.Query()
	tokenString := vars["token"][0]
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("hmacSampleSecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["exp"])
	} else {
		fmt.Println("error: ", err)
	}
	rb := dto.ResBody{
		Code: 0,
		Msg:  "登陆成功",
	}
	rb.ReturnToFE(w)
}
