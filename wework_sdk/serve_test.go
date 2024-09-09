package wework_sdk

import (
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestServe(t *testing.T) {

	http.HandleFunc("/push_url", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		pahts := strings.Split(path, "/")
		recevierId := ""
		if pahts[1] != "" {
			recevierId = pahts[1]
		} else {
			log.Println("path:", path)
			writer.Write([]byte("success"))
			return
		}
		client := New(recevierId, &Config{
			EncodeAesKey: "jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C",
			Token:        "QDG6eK",
		})
		if request.Method == "POST" {
			//client.DecryptMsg(request.FormValue("msg_signature))
		}
		if request.Method == "GET" {
			resp, err := client.VerifyURL(request.FormValue("msg_signature"), request.FormValue("timestamp"), request.FormValue("nonce"), request.FormValue("echostr"))
			if err != nil {
				writer.Write([]byte(err.Error()))
			} else {
				writer.Write([]byte(resp))
			}
		}
	}) //      设置访问路由

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("success"))
	//})
	log.Fatal(http.ListenAndServe(":8893", nil))
}
