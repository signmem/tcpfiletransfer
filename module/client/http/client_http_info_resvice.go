package http

import (
	"fmt"
	"github.com/signmem/tcpfiletransfer/module/client/send"
	"net/http"
	"github.com/signmem/tcpfiletransfer/g"
)

func httpFileUploadRevice() {
	http.HandleFunc("/api/v1/upload",
		func(w http.ResponseWriter, r *http.Request) {

			clientIP, err := g.GetClientIP(r)
			if err != nil {
				clientIP = "None"
			}


			fileInfo, err := g.HTTPCheckContent(r)

			if err != nil {
				msg := fmt.Sprintf("httpFileGet() client: %s, %s", clientIP, err)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)

				return
			}

			fileInfo.FSClient = clientIP

			msg := fmt.Sprintf("HttpFileUploadRevice() %s", fileInfo.String())
			g.Logger.Debug(msg)

			send.ReveFromHTTPToTcp(fileInfo)
			w.Write([]byte(msg))
			return
		})
}