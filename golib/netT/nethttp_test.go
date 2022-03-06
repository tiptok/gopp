package netT

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCommMethod(t *testing.T) {
	contentType := http.DetectContentType([]byte("{\"name\":\"123\"}"))
	log.Printf("Detect type:%s", contentType)

	maj, min, err := http.ParseHTTPVersion("HTTP/1.1")
	if !err {
		log.Printf("Parse version error.")
	}
	log.Printf("Version:%v.%v", maj, min)
}

// client do request
func TestRequest(t *testing.T) {
	client := &http.Client{}
	url := "http://dis.dfmwl.com/api/GpsIncome"
	request, err := http.NewRequest("POST", url, strings.NewReader(`[{"truck_num":"渝BY6990","update_time":"2019-02-25 09:54:40","province":"广东省","city":"广州市","road":"南沙区黄阁大道中 海港酒家(海头新街)东北241米","speed":"0.0","lng":106.264632,"lat":29.400716,"source":"赛格"}]`))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("request error", url, err)
	}
	rsp, err := client.Do(request)
	ReadResponse(rsp, err)
}

func ReadResponse(rsp *http.Response, err error) {
	if err != nil {
		log.Println("Request error.", err)
	}
	log.Println(rsp.Proto)
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		log.Println("Read Response error.", err)
	}
	log.Println(rsp.StatusCode)
	log.Println("ContentLen:", rsp.ContentLength)
	log.Println("Header:", rsp.Header)
	log.Println(string(body))
}

//http post
func TestHttpPost(t *testing.T) {
	url := "http://dis.dfmwl.com/api/GpsIncome"
	rsp, err := http.Post(url, "application/json", strings.NewReader(`[{"truck_num":"渝BY6990","update_time":"2019-02-25 09:54:40","province":"广东省","city":"广州市","road":"南沙区黄阁大道中 海港酒家(海头新街)东北241米","speed":"0.0","lng":106.264632,"lat":29.400716,"source":"赛格"}]`))
	ReadResponse(rsp, err)
}

//http get
func TestHttpGet(t *testing.T) {
	url := "http://dis.dfmwl.com/api/GpsIncome"
	rsp, err := http.Get(url)
	ReadResponse(rsp, err)
}
