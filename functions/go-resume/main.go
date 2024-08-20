package handler

import (
	"embed"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

//go:embed static/*
var embedReader embed.FS

func Main(Context openruntimes.Context) openruntimes.Response {
	if Context.Req.Method == "GET" {

		resumeHtml, _ := embedReader.ReadFile(("static/resume.html"))

		return Context.Res.Text(string(resumeHtml),
			Context.Res.WithStatusCode(200),
			Context.Res.WithHeaders(map[string]string{
				"Content-Type": "text/html",
			}))
	}
	return Context.Res.Text("Bad request", Context.Res.WithStatusCode(404))
}
