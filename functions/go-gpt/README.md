# Go GPT

Ask question, and let OpenAI GPT-3.5-turbo answer.

## 🧰 Usage

### POST /

Query the model for a completion.

**Parameters**

| Name         | Description                          | Location | Type               | Sample Value                  |
| ------------ | ------------------------------------ | -------- | ------------------ | ----------------------------- |
| Content-Type | The content type of the request body | Header   | `application/json` | N/A                           |
| prompt       | Text to prompt the model             | Body     | String             | `Write a haiku about Mondays` |

Sample `200` Response:

Response from the model.

```json
{
  "ok": true,
  "completion": "Monday's heavy weight, Dawning with a sigh of grey, Hopeful hearts await."
}
```

Sample `400` Response:

Response when the wrong HTTP action is used.

```json
{
  "ok": false,
  "error": "Bad request"
}
```

Sample `500` Response:

Response when the model fails to respond.

```json
{
  "ok": false,
  "error": "Failed to query model."
}
```

## ⚙️ Configuration

| Setting           | Value         |
| ----------------- | ------------- |
| Runtime           | Go (1.22)     |
| Entrypoint        | `main.go`     |
| Permissions       | `any`         |
| Timeout (Seconds) | 15            |

## 🔒 Environment Variables

### OPENAI_API_KEY

A unique key used to authenticate with the OpenAI API. Please note that this is a paid service and you will be charged for each request made to the API. For more information, see the [OpenAI pricing page](https://openai.com/pricing/).

| Question      | Answer                                                                      |
| ------------- | --------------------------------------------------------------------------- |
| Required      | Yes                                                                         |
| Sample Value  | `sk-wzG...vcy`                                                              |
| Documentation | [OpenAI Docs](https://platform.openai.com/docs/quickstart/add-your-api-key) |
