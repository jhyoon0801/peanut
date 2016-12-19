package util

import (
	"sync/atomic"
	//	"unsafe"
)

var baseOfUniqueId *int64 = new(int64)

// In go, Abstract class without abstract method can't create.
// So this struct AbstractConstant just has name same with java's AbstractConstant
type AbstractConstant struct {
	id         int
	name       string
	uniquifier int64
}

func (ac *AbstractConstant) Id() int {
	return ac.id
}

func (ac *AbstractConstant) Name() string {
	return ac.name
}

func (ac *AbstractConstant) ToString() string {
	return ac.name
}

func NewAbstractConstant(id int, name string) *AbstractConstant {
	return &AbstractConstant{id, name, atomic.AddInt64(baseOfUniqueId, 1)}
}

//func (ac *abstractConstant) equal(obj unsafe.Pointer) bool {
//	return ac == Pointer
//}

//func (ac *abstractConstant) compareTo(obj unsafe.Pointer) int {
//	if ac == Pointer {
//		return 0
//	}

//	obj

//	return
//}

/* Todo */
/*
	Are these function require?
		- hashCode
		-
*/
