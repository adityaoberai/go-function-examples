package handler

import (
	"openruntimes/handler/services"
	"os"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

type RequestBody struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}

type ResponseBody struct {
	LongUrl string `json:"longUrl"`
}

func Main(Context openruntimes.Context) openruntimes.Response {
	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_FUNCTION_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_FUNCTION_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)

	databases := appwrite.NewDatabases(client)

	dbId := os.Getenv("APPWRITE_DB_ID")
	collId := os.Getenv("APPWRITE_COLL_ID")

	services.InitialiseDatabase(Context, *databases, dbId, collId)

	if Context.Req.Method == "POST" {
		var requestBody RequestBody
		err := Context.Req.BodyJson(&requestBody)
		if err != nil {
			Context.Error(err)
			return Context.Res.Json(map[string]interface{}{
				"ok":    false,
				"error": "Missing request body",
			}, Context.Res.WithStatusCode(400))
		}

		document, err := databases.CreateDocument(
			dbId,
			collId,
			requestBody.ShortUrl,
			map[string]interface{}{
				"longUrl": requestBody.LongUrl,
			},
		)

		if err != nil {
			Context.Error(err)
			return Context.Res.Json(map[string]interface{}{
				"ok":    false,
				"error": "Failed to create shortened URL",
			}, Context.Res.WithStatusCode(500))
		}

		return Context.Res.Json(map[string]interface{}{
			"ok":       true,
			"response": document,
		}, Context.Res.WithStatusCode(200))
	}

	if Context.Req.Method == "GET" {
		path := Context.Req.Path
		if path == "/" {
			return Context.Res.Text("Welcome to the URL shortener service\n\nAdd a short URL to the path to redirect to the long URL\n", Context.Res.WithStatusCode(200))
		}

		shortUrl := path[1:]

		document, err := databases.GetDocument(dbId, collId, shortUrl)

		if err != nil {
			Context.Error(err)
			return Context.Res.Text("URL not found", Context.Res.WithStatusCode(400))
		}

		var responseBody ResponseBody
		document.Decode(&responseBody)

		return Context.Res.Redirect(responseBody.LongUrl, Context.Res.WithStatusCode(301))
	}
	return Context.Res.Empty()
}
