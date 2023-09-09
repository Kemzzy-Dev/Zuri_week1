// implements the api as a serverless function on vercel
package handle

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResObj struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")
	now := time.Now().UTC()
	day := now.Weekday().String()
	githubFile := "https://github.com/joey1123455/nzuri_server_week1/blob/main/api/index.go"
	githubRepo := "https://github.com/joey1123455/nzuri_server_week1"
	statusCode := 200
	res := &ResObj{
		SlackName:     slackName,
		CurrentDay:    day,
		UtcTime:       now.Format("2006-01-02T15:04:05Z"),
		Track:         track,
		GithubFileUrl: githubFile,
		GithubRepoUrl: githubRepo,
		StatusCode:    statusCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
