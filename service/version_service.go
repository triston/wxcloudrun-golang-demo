package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// VersionResult 返回结构
type VersionResult struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"errorMsg,omitempty"`
	Version  string `json:"version"`
}

// VersionHandler 版本接口
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	res := &VersionResult{}

	//  打印请求头
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("Header Key: " + k + " value: " + vv)
		}
	}

	getUserPhone()

	if r.Method == http.MethodGet {
		res.Version = "20220813.01"
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func getUserPhone() {
	resp, err := http.Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber")
	fmt.Println("resp", resp)
}
