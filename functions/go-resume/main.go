package handler

import (
	"openruntimes/handler/services"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

func Main(Context openruntimes.Context) openruntimes.Response {
	if Context.Req.Method == "GET" {
		resumeHtml := services.GetStaticFile(Context, "resume.html")

		return Context.Res.Text(resumeHtml, Context.Res.WithStatusCode(200), Context.Res.WithHeaders(map[string]string{
			"Content-Type": "text/html; charset=utf-8",
		}))
	}
	return Context.Res.Text("Bad request", Context.Res.WithStatusCode(400))
}
