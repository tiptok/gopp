package htmlT

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"testing"
)

type Vehicle struct {
	PlateNum string
	SimNum   string
	Enable   bool
}

func TestTemplate1(t *testing.T) {
	v := Vehicle{PlateNum: "闽A12345", SimNum: "18860152301"}
	tp := template.New("template demo")
	tp, _ = tp.Parse("car:{{.PlateNum}} simnum:{{.SimNum}}")
	err := tp.Execute(os.Stdout, v)
	if err != nil {
		log.Fatal(err)
	}
}

const tpstr = `
<h1>Vehicle</h1>
<ul>
    {{range .Vehicle}}
        {{if .Enable}}
            <li>{{.PlateNum}} simnum:{{.SimNum}}</li>
        {{else}}
            <li><s>car:{{.PlateNum}} simnum:{{.SimNum}}</s></li>
        {{end}}
    {{end}}
</ul>
`

func TestTpParseFile(t *testing.T) {
	vs := []Vehicle{
		{PlateNum: "闽A12345", SimNum: "18860152301", Enable: true},
		{PlateNum: "闽A15432", SimNum: "18860152321", Enable: true},
		{PlateNum: "闽A16432", SimNum: "18860152321", Enable: false},
	}
	tp, err := template.New("CarList").Parse(tpstr)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tp.Execute(writer, struct {
			Vehicle []Vehicle
		}{vs})
	})
	http.ListenAndServe(":8080", nil)
}
