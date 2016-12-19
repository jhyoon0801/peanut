package channel

import (
	"errors"
	"peanut/util"
	"reflect"
)

// TODO
/*
	1. valueOf(Class<?> firstNameComponent, String secondNameComponent)
*/

const (
	INT_TYPE  int  = 1
	BOOL_TYPE bool = true
)

type ChannelOptionConstantPool struct {
	*util.AbstractConstantPool
}

func newConstant(id int, name string) util.Constant {
	return &ChannelOption{util.NewAbstractConstant(id, name), nil}
}

var pool *ChannelOptionConstantPool = &ChannelOptionConstantPool{util.NewAbstractConstantPool(newConstant)}

type ChannelOption struct {
	*util.AbstractConstant
	t reflect.Type
}

func ValueOf(name string, t reflect.Type) *ChannelOption {
	var channelOption *ChannelOption = pool.ValueOf(name).(*ChannelOption)
	channelOption.t = t
	return channelOption
}

func Exists(name string) (bool, error) {
	return pool.Exists(name)
}

func NewInstance(name string, t reflect.Type) *ChannelOption {
	var newChannelOption *ChannelOption = pool.NewInstance(name).(*ChannelOption)
	newChannelOption.t = t
	return newChannelOption
}

func Validate(value interface{}) error {
	if nil == value {
		return errors.New("value is nil")
	}
	return nil
}

// var ALLOCATOR *ChannelOption = valueOf("ALLOCATOR", ByteBufAllocator)
//var RCVBUF_ALLOCATOR *ChannelOption = valueOf("RCVBUF_ALLOCATOR", RecvByteBufAllocator)
//var MESSAGE_SIZE_ESTIMATOR *ChannelOption = valueOf("MESSAGE_SIZE_ESTIMATOR", MessageSizeEstimator)

var CONNECT_TIMEOUT_MILLIS *ChannelOption = ValueOf("CONNECT_TIMEOUT_MILLIS", reflect.TypeOf(INT_TYPE))

// TODO MAX_MESSAGES_PER_READ will be removed
var CHANNEL_OPTION_MAX_MESSAGES_PER_READ *ChannelOption = ValueOf("MAX_MESSAGES_PER_READ", reflect.TypeOf(INT_TYPE))
var CHANNEL_OPTION_WRITE_SPIN_COUNT *ChannelOption = ValueOf("WRITE_SPIN_COUNT", reflect.TypeOf(INT_TYPE))

// TODO WRITE_BUFFER_HIGH_WATER_MARK will be removed
var WRITE_BUFFER_HIGH_WATER_MARK *ChannelOption = ValueOf("WRITE_BUFFER_HIGH_WATER_MARK", reflect.TypeOf(INT_TYPE))

// TODO WRITE_BUFFER_LOW_WATER_MARK will be removed
var WRITE_BUFFER_LOW_WATER_MARK *ChannelOption = ValueOf("WRITE_BUFFER_LOW_WATER_MARK", reflect.TypeOf(INT_TYPE))

//var WRITE_BUFFER_WATER_MARK *ChannelOption =  valueOf("WRITE_BUFFER_WATER_MARK", WriteBufferWaterMark)
var ALLOW_HALF_CLOSURE *ChannelOption = ValueOf("ALLOW_HALF_CLOSURE", reflect.TypeOf(BOOL_TYPE))
var AUTO_READ *ChannelOption = ValueOf("AUTO_READ", reflect.TypeOf(BOOL_TYPE))

// TODO AUTO_CLOSE will be removed
var AUTO_CLOSE *ChannelOption = ValueOf("AUTO_CLOSE", reflect.TypeOf(BOOL_TYPE))

var SO_BROADCAST *ChannelOption = ValueOf("SO_BROADCAST", reflect.TypeOf(BOOL_TYPE))
var SO_KEEPALIVE *ChannelOption = ValueOf("SO_KEEPALIVE", reflect.TypeOf(BOOL_TYPE))
var SO_SNDBUF *ChannelOption = ValueOf("SO_SNDBUF", reflect.TypeOf(INT_TYPE))
var SO_RCVBUF *ChannelOption = ValueOf("SO_RCVBUF", reflect.TypeOf(INT_TYPE))
var SO_REUSEADDR *ChannelOption = ValueOf("SO_REUSEADDR", reflect.TypeOf(BOOL_TYPE))
var SO_LINGER *ChannelOption = ValueOf("SO_LINGER", reflect.TypeOf(INT_TYPE))
var SO_BACKLOG *ChannelOption = ValueOf("SO_BACKLOG", reflect.TypeOf(INT_TYPE))
var SO_TIMEOUT *ChannelOption = ValueOf("SO_TIMEOUT", reflect.TypeOf(INT_TYPE))

var IP_TOS *ChannelOption = ValueOf("IP_TOS", reflect.TypeOf(INT_TYPE))

//var IP_MULTICAST_ADDR *ChannelOption = valueOf("IP_MULTICAST_ADDR", int)
//var IP_MULTICAST_IF *ChannelOption = valueOf("IP_MULTICAST_IF", int)
var IP_MULTICAST_TTL *ChannelOption = ValueOf("IP_MULTICAST_TTL", reflect.TypeOf(INT_TYPE))
var IP_MULTICAST_LOOP_DISABLED *ChannelOption = ValueOf("IP_MULTICAST_LOOP_DISABLED", reflect.TypeOf(BOOL_TYPE))

var TCP_NODELAY *ChannelOption = ValueOf("TCP_NODELAY", reflect.TypeOf(BOOL_TYPE))

// TODO DATAGRAM_CHANNEL_ACTIVE_ON_REGISTRATION will be removed
var DATAGRAM_CHANNEL_ACTIVE_ON_REGISTRATION *ChannelOption = ValueOf("DATAGRAM_CHANNEL_ACTIVE_ON_REGISTRATION", reflect.TypeOf(BOOL_TYPE))

var SINGLE_EVENTEXECUTOR_PER_GROUP *ChannelOption = ValueOf("SINGLE_EVENTEXECUTOR_PER_GROUP", reflect.TypeOf(BOOL_TYPE))
