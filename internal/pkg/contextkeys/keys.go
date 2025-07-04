package contextkeys

type contextKey string

const (
	SessionKey contextKey = "session"
	ToastKey   contextKey = "toast"
	HostKey    contextKey = "host"
)
