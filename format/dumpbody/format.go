package dumpbody

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/motemen/go-loghttp"
	_ "github.com/motemen/go-loghttp/global"
)

func init() {
	loghttp.DefaultLogRequest = func(req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
		dump, _ := httputil.DumpRequest(req, false)
		log.Printf("---> [Req] %s\n", string(dump))
		log.Printf("---> [Req Body] %s\n\n", body)
	}

	loghttp.DefaultLogResponse = func(resp *http.Response) {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		resp.Body = ioutil.NopCloser(bytes.NewReader(body))
		dump, _ := httputil.DumpResponse(resp, false)
		log.Printf("<--- [Resp] %s\n", string(dump))
		log.Printf("<--- [Resp Body] %s\n\n", body)
	}
}
