package util

import (
	"time"
)

type Timer interface {
	NewTimeout(tt TimerTask, delay time.Duration, unit time.Duration) Timeout
	Stop() []Timeout
}
