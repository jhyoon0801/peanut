package util

import (
	"reflect"
)

type ConstantPool_AttributeKey struct {
	*AbstractConstantPool
}

func newConstant_AttributeKey(id int, name string) Constant {
	return &AttributeKey{NewAbstractConstant(id, name), nil}
}

var attributeKeyPool *ConstantPool_AttributeKey = &ConstantPool_AttributeKey{NewAbstractConstantPool(newConstant_AttributeKey)}

type AttributeKey struct {
	*AbstractConstant
	t reflect.Type
}

func Exists(name string) (bool, error) {
	return attributeKeyPool.Exists(name)
}

func NewInstance(name string, t reflect.Type) *AttributeKey {
	var newAttributeKey *AttributeKey = attributeKeyPool.NewInstance(name).(*AttributeKey)
	newAttributeKey.t = t
	return newAttributeKey
}

func ValueOf_AttributeKey(name string, t reflect.Type) *AttributeKey {
	attributeKey := attributeKeyPool.ValueOf(name).(*AttributeKey)
	attributeKey.t = t
	return attributeKey
}
