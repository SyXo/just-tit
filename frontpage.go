package main

import (
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"html/template"
)

func frontpage() events.APIGatewayProxyResponse {
	headers := map[string]string{
		"Content-Type":  "text/html; charset=utf-8",
		"Cache-Control": "max-age=31536000",
	}

	// Build HTML from template
	web := template.Must(template.New("frontpage").Funcs(TemplateFunctions).ParseFiles(
		"html/"+Theme+"/template.html",
		"html/"+Theme+"/video/container.html",
		"html/"+Theme+"/search/container.html",
		"html/"+Theme+"/search/pornhub.html",
		"html/"+Theme+"/search/redtube.html",
		"html/"+Theme+"/search/tube8.html",
		"html/"+Theme+"/search/youporn.html",
	))

	replace := TemplateData{
		PageTitle:    "Just Tit",
		PageMetaDesc: "The most optimized adult video search engine",
	}

	// Build result divs
	var buff bytes.Buffer
	web.ExecuteTemplate(&buff, "layout", replace)
	body := buff.String()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       body,
	}
}
