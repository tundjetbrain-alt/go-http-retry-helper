package retry

import (
	"net/http"
	"time"
)

func BackoffDelay(attempt int) time.Duration {
	if attempt <= 0 {
		return 100 * time.Millisecond
	}
	return time.Duration(1<<uint(attempt)) * 100 * time.Millisecond
}

var DefaultClient = &http.Client{Timeout: 10 * time.Second}
