package dumpheader

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/motemen/go-loghttp"
	_ "github.com/motemen/go-loghttp/global"
)

func init() {
	loghttp.DefaultLogRequest = func(req *http.Request) {
		dump, _ := httputil.DumpRequest(req, false)
		log.Printf("---> [Req] %s\n", string(dump))
	}

	loghttp.DefaultLogResponse = func(resp *http.Response) {
		dump, _ := httputil.DumpResponse(resp, false)
		log.Printf("<--- [Resp] %s\n", string(dump))
	}
}
