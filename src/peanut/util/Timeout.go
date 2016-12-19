package util

type Timeout interface {
	Timer() Timer
	Task() TimerTask
	IsExpired() bool
	IsCancelled() bool
	Cancel() bool
}
