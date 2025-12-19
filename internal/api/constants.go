package api

import "time"

const (
	// DefaultUserID is the default user ID when none is provided.
	DefaultUserID = "default"

	// ExecuteTimeout is the timeout for code execution requests.
	ExecuteTimeout = 15 * time.Second
)
