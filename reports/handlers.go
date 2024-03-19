package reports

import (
	"fmt"
	"html/template"
	"net/http"
)

func ReportsMainHandler(w http.ResponseWriter, r *http.Request) {
	reports := IndexReportService()
	fmt.Println(reports)
	tmpl, _ := template.ParseFiles("./reports/templates/index.html")
	tmpl.Execute(w, reports)
}
