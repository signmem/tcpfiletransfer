package http

import (
	"net/http"
	"encoding/json"
	"github.com/signmem/tcpfiletransfer/g"
)

func init() {
	healthCheck()
	httpFileUploadRevice()
}

type Dto struct {
	Msg	string		`json:"msg"`
	Data    interface{}     `json:"data"`
}


func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Data: data})
}

func RenderMsgJson(w http.ResponseWriter, msg string) {
	RenderJson(w, map[string]string{"msg": msg})
}

func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error())
		return
	}

	RenderDataJson(w, data)
}

func Start() {

	address := g.Config().Http.Address
	port := g.Config().Http.Port
	listen := address + ":" + port

	s := &http.Server{
		Addr:           listen,
		MaxHeaderBytes: 1 << 30,
	}

	g.Logger.Infof("listening %s", listen)
	g.Logger.Fatalln(s.ListenAndServe())
}

func healthCheck() {
	http.HandleFunc("/health_check",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
}
