package types

type Response struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}
