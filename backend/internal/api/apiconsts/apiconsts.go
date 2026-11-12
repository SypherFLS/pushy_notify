package apiconsts

type ContextKey string

const (
	GithubEvents    ContextKey = "X-GitHub-Event"
	GithubDelivery  ContextKey = "X-GitHub-Delivery"
	GithubSignature ContextKey = "X-Hub-Signature-256"
)
