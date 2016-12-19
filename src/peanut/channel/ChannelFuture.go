package channel

import (
	"peanut/util/concurrent"
)

type ChannelFuture struct {
	*concurrent.Future
	*concurrent.PeanutFuture
}
