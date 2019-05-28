package dto

import (
	"encoding/json"
	"net/http"
)

// ResBody 返回体
type ResBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"errMsg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// ReturnToFE 返回前端数据
func (rb *ResBody) ReturnToFE(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	res, _ := json.Marshal(rb)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
