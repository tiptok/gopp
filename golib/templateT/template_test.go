package templateT

import (
	"bytes"
	"html/template"
	"testing"
)

func TestCompare(t *testing.T) {
	temp := `
{{ if eq .arg "Delete" "Remove"}}
this is delete/remove item
{{ else if eq .arg "Create"}}
this is create item
{{ else }}
this is update/list/save item
{{ end }}
`
	tp, err := template.New("xxx").Parse(temp)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	tp.Execute(buf, map[string]interface{}{"arg": "Create"})
	t.Log(buf.String())
}

func TestIterate(t *testing.T) {
	type Item struct {
		Name  string
		Price int
	}
	type ViewData struct {
		Name  string
		Items []Item
	}
	viewData := ViewData{Name: "view1", Items: []Item{{Name: "西瓜", Price: 10}, {Name: "香蕉", Price: 5}, {Name: "橘子", Price: 8}}}
	temp := `
{{range .Items}}
  <div class="item">
    <h3 class="name">{{.Name}}</h3>
    <span class="price">{{.Price}}</span>
  </div>
{{end}}
`
	tp, err := template.New("xxx").Parse(temp)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	//tp.Execute(buf,map[string]interface{}{"Items":viewData.Items}

	tp.Execute(buf, viewData)
	t.Log(buf.String())
}

func TestRangeFileds(t *testing.T) {
	type item struct {
		Name   string
		Column string
		Type   string
	}
	var fileds = []item{}
	fileds = append(fileds, item{Name: "Name", Column: "name", Type: "string"}, item{Name: "Time", Column: "time", Type: "time.Time"})

	temp := `
{{range .Fields}}	if v, ok := data["{{.Column}}"]; ok {
		m.{{.Name}} = v.({{.Type}})
	}
{{end}}
`
	tp, err := template.New("xxx").Parse(temp)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	//tp.Execute(buf,map[string]interface{}{"Items":viewData.Items}

	tp.Execute(buf, map[string]interface{}{"Fields": fileds})
	t.Log(buf.String())
}
