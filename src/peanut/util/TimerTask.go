package util

type TimerTask interface {
	run(to Timeout) error
}
