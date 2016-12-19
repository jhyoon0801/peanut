package util

import (
	"errors"
	"reflect"
)

// TODO
/*
	1. valueOf(Class<?> firstNameComponent, String secondNameComponent)
	2. initCause
	3. fillInStackTrace
	4. hashCode
	5. compareTo
*/

// Implements Constant and error
type ConstantPool_Signal struct {
	*AbstractConstantPool
}

func newConstant_Signal(id int, name string) Constant {
	return &Signal{&signalConstant{NewAbstractConstant(id, name)}}
}

var signalPool *ConstantPool_Signal = &ConstantPool_Signal{NewAbstractConstantPool(newConstant_Signal)}

func ValueOf_Signal(name string) *Signal {
	return signalPool.ValueOf(name).(*Signal)
}

type Signal struct {
	constant *signalConstant
}

func (s *Signal) Error() string {
	return s.Name()
}

func (s *Signal) Expect(signal *Signal) error {
	if s == signal {
		return nil
	}
	return errors.New("unexpected signal: " + signal.ToString())
}

func (s *Signal) Id() int {
	return s.constant.Id()
}

func (s *Signal) Name() string {
	return s.constant.Name()
}

func (s *Signal) ToString() string {
	return s.constant.ToString()
}

func (s *Signal) Equal(obj interface{}) bool {
	if reflect.TypeOf(s) != reflect.TypeOf(obj) {
		return false
	}
	return s == obj.(*Signal)
}

func (s *Signal) CompareTo(signal *Signal) int {
	if s == signal {
		return 0
	}
	return 1
}

type signalConstant struct {
	*AbstractConstant
}
