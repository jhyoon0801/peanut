package util

import (
	"errors"
	"sync/atomic"
)

// TODO
/*
	1. valueOf(Class<?> firstNameComponent, String secondNameComponent)
*/

type AbstractConstantPool struct {
	newConstant func(int, string) Constant
	constants   map[string]Constant
	nextId      int32
}

func NewAbstractConstantPool(newConstant func(int, string) Constant) *AbstractConstantPool {
	rt := new(AbstractConstantPool)
	rt.newConstant = newConstant
	rt.constants = make(map[string]Constant)
	rt.nextId = 1
	return rt
}

func (cp *AbstractConstantPool) ValueOf(name string) Constant {
	if err := cp.checkNotNullAndNotEmpty(name); nil != err {
		panic(err.Error())
	}
	return cp.getOrCreate(name)
}

func (cp *AbstractConstantPool) Exists(name string) (bool, error) {
	if err := cp.checkNotNullAndNotEmpty(name); nil != err {
		return false, err
	}
	_, exist := cp.constants[name]
	return exist, nil
}

func (cp *AbstractConstantPool) NewInstance(name string) Constant {
	if err := cp.checkNotNullAndNotEmpty(name); nil != err {
		panic(err.Error())
	}
	return cp.create(name)
}

func (cp *AbstractConstantPool) checkNotNullAndNotEmpty(name string) error {
	if 0 == len(name) {
		return errors.New("empty name")
	}
	return nil
}

func (cp *AbstractConstantPool) getOrCreate(name string) Constant {
	cst := cp.constants[name]
	if nil == cst {
		cst = cp.newConstant(int(atomic.AddInt32(&cp.nextId, 1)), name)
		cp.constants[name] = cst
	}
	return cst
}

func (cp *AbstractConstantPool) create(name string) Constant {
	cst := cp.constants[name]
	if nil != cst {
		panic(name + " is already in use")
	}

	cst = cp.newConstant(int(atomic.AddInt32(&cp.nextId, 1)), name)
	cp.constants[name] = cst
	return cst
}
