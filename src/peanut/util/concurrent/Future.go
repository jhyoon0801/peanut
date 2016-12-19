package concurrent

import (
	"time"
)

type Future interface {
	Cancel(bool) bool
	IsCancelled() bool
	IsDone() bool
	Get() interface{}
	GetWithTimeout() interface{}
}
