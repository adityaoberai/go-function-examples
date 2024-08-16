# Go Resume

Share your resume as an HTML webpage.

## 🧰 Usage

### GET /

- Returns an HTML webpage that renders in your browser

### GET, POST, PUT, PATCH, DELETE /

- Returns a "Bad request" text message

**Response**

Sample `400` Response:

```
Bad request
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
