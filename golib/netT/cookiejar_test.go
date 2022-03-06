package netT

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func InitAll() {
	gCurCookies = nil
	gCurCookieJar, _ = cookiejar.New(nil)
}

func getUrlRespHtml(url string) string {
	log.Printf("getUrlRespHtml, url=%s\n", url)

	//var sRspHtml string = ""
	httpCli := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	httpReq, err := http.NewRequest("GET", url, nil)

	httpResp, err := httpCli.Do(httpReq)
	if err != nil {
		log.Panicf("http get url=%s response error=%s\n", url, err.Error())
	}
	log.Printf("httpResp.Header=%s\n", httpResp.Header)
	log.Printf("httpResp.Status=%s\n", httpResp.Status)
	body, _ := ioutil.ReadAll(httpResp.Body)
	//log.Printf("httpResp.Status=%s\n", string(body))
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)
	return string(body)
}

func dbgPrintCurCookies() {
	var icNum int = len(gCurCookies)
	log.Println("cookieNum=", icNum)
	for i := 0; i < icNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		log.Printf("Name\t=%s\n", curCk.Name)
		log.Printf("Value\t=%s\n", curCk.Value)
		log.Printf("Path\t=%s\n", curCk.Path)
		log.Printf("Domain\t=%s\n", curCk.Domain)
		log.Printf("Expires\t=%s\n", curCk.Expires)
		log.Printf("RawExpires\t=%s\n", curCk.RawExpires)
		log.Printf("MaxAge\t=%d\n", curCk.MaxAge)
		log.Printf("Secure\t=%t\n", curCk.Secure)
		log.Printf("HttpOnly\t=%t\n", curCk.HttpOnly)
		log.Printf("Raw\t=%s\n", curCk.Raw)
		log.Printf("Unparsed\t=%s\n", curCk.Unparsed)
	}
}

func TestCookies(t *testing.T) {
	InitAll()

	log.Println("====== 步骤1：获得BAIDUID的Cookie ======")
	var sBD string = "http://www.baidu.com/"
	_ = getUrlRespHtml(sBD)
	//log.Println(respHtml)
	dbgPrintCurCookies()

	log.Println("====== 步骤2：提取login_token ======")

}
