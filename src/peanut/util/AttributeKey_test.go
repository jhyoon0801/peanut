package util

import (
	"reflect"
	"strings"
	"testing"
)

type structForTest1 struct {
}

type structForTest2 struct {
}

func TestValueOf_AttributeKey(t *testing.T) {
	attributeKey_1 := ValueOf_AttributeKey("TEST_ATTRIBUTEKEY", reflect.TypeOf(1))
	attributeKey_2 := ValueOf_AttributeKey("TEST_ATTRIBUTEKEY", reflect.TypeOf(1))

	if !strings.EqualFold(attributeKey_1.Name(), "TEST_ATTRIBUTEKEY") ||
		!strings.EqualFold(attributeKey_2.Name(), "TEST_ATTRIBUTEKEY") ||
		attributeKey_1.Id() != attributeKey_2.Id() ||
		attributeKey_1.t.Kind() != attributeKey_2.t.Kind() {
		t.Errorf("Test TestValueOf_AttributeKey has failed")
	}

	attributeKey_1 = ValueOf_AttributeKey("TEST_ATTRIBUTEKEY1", reflect.TypeOf(structForTest1{}))
	attributeKey_2 = ValueOf_AttributeKey("TEST_ATTRIBUTEKEY2", reflect.TypeOf(structForTest2{}))

	if !strings.EqualFold(attributeKey_1.Name(), "TEST_ATTRIBUTEKEY1") ||
		!strings.EqualFold(attributeKey_2.Name(), "TEST_ATTRIBUTEKEY2") ||
		attributeKey_1.t == attributeKey_2.t {
		t.Errorf("Test TestValueOf_AttributeKey has failed")
	}
}

func TestNewInstance(t *testing.T) {
	attributeKey := NewInstance("TestNewInstance", reflect.TypeOf(""))
	if !strings.EqualFold(attributeKey.Name(), "TestNewInstance") ||
		reflect.String != attributeKey.t.Kind() {
		t.Errorf("Test NewInstance has failed")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test NewInstance has sucess")
		}
	}()
	attributeKey = NewInstance("TestNewInstance", reflect.TypeOf(""))

	t.Errorf("Test NewInstance has failed")
}

func TestExists(t *testing.T) {
	ValueOf_AttributeKey("TestNewInstance", reflect.TypeOf(1.0))
	if exist, _ := Exists("TestNewInstance"); false == exist {
		t.Errorf("Test TestExists has failed. [exist : %v]", exist)
	}
	if exist, err := Exists("TEST_INVALID_ATTRIBUTEKEY"); true == exist || nil != err {
		t.Errorf("Test TestExists has failed [exist : %v][err : %v]", exist, err)
	}
	if exist, err := Exists(""); true == exist || nil == err {
		t.Errorf("Test TestExists has failed [exist : %v][err : %v]", exist, err)
	}
}
