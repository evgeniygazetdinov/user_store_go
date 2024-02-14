package reports

import (
	"html/template"
	"net/http"
)

func ReportsMainHandler(w http.ResponseWriter, r *http.Request) {
	reports := IndexReportService()
	tmpl, _ := template.ParseFiles("./reports/templates/index.html")
	tmpl.Execute(w, reports)
}
