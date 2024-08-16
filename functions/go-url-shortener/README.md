# Go URL Shortner
Generate URL with short ID and redirect to the original URL when visited.

## 🧰 Usage

### GET /:shortId

Redirects to shortId's original URL.

**Parameters**

| Name    | Description                      | Location | Type   | Sample Value |
| ------- | -------------------------------- | -------- | ------ | ------------ |
| shortId | Short ID to lookup original URL. | Path     | String | `appwrite`    |

**Response**

Sample `302` Response:

Redirects to the original URL.

```text
Location: https://appwrite.io
```

Sample `404` Response:

When no URL is found for the short ID.

```text
URL not found
```

### POST /

Create a new short ID for a URL.

**Parameters**

| Name         | Description                                           | Location | Type               | Sample Value                                                   |
| ------------ | ----------------------------------------------------- | -------- | ------------------ | -------------------------------------------------------------- |
| Content-Type | Content type                                          | Header   | `application/json` |
| longUrl      | Long URL to shorten                                   | Body     | String             | `https://mywebapp.com/pages/hugelongurl?with=query&params=123` |
| shortId      | Short ID to use                                       | Body     | String             | `discord`                                                      |

**Response**

Sample `200` Response:

Returns the short URL and the original URL.

```json
{
  "ok": true,
  "shortId": "appwrite",
  "longUrl": "https://appwrite.io"
}
```

Sample `400` Response:

When the request body doesn't deserialise properly.

```json
{
  "ok": false,
  "error": "Missing request body"
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

No environment variables required.
