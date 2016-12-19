package channel

import (
	"reflect"
	"strings"
	"testing"
)

func TestExists(t *testing.T) {
	if exist, _ := Exists("SO_BACKLOG"); false == exist {
		t.Errorf("Test Exists : SO_BACKLOG. [exist : %v]")
	}

	if exist, _ := Exists("ALLOW_HALF_CLOSURE"); false == exist {
		t.Errorf("Test Exists : ALLOW_HALF_CLOSURE. [exist : %v]")
	}

	if exist, _ := Exists("SO_TIMEOUT"); false == exist {
		t.Errorf("Test Exists : SO_TIMEOUT. [exist : %v]")
	}

	if exist, err := Exists("TEST_INVALID_NAME"); true == exist || nil != err {
		t.Errorf("Test Exists : TEST_INVALID_NAME. [exist : %v][err : %v]", exist, err)
	}

	if exist, err := Exists(""); true == exist || nil == err {
		t.Errorf("Test Exists : . [exist : %v][err : %v]", exist, err)
	}
}

func TestValueOf(t *testing.T) {
	channelOption_1 := ValueOf("SO_BACKLOG", reflect.TypeOf(INT_TYPE))
	channelOption_2 := ValueOf("SO_BACKLOG", reflect.TypeOf(INT_TYPE))
	if !strings.EqualFold(channelOption_1.Name(), "SO_BACKLOG") ||
		!strings.EqualFold(channelOption_2.Name(), "SO_BACKLOG") ||
		reflect.Int != channelOption_1.t.Kind() ||
		channelOption_1.Id() != channelOption_2.Id() {
		t.Errorf("Test ValueOf has failed")
	}

	channelOption := ValueOf("ALLOW_HALF_CLOSURE", reflect.TypeOf(BOOL_TYPE))
	if !strings.EqualFold(channelOption.Name(), "ALLOW_HALF_CLOSURE") ||
		reflect.Bool != channelOption.t.Kind() {
		t.Errorf("Test ValueOf has failed")
	}
}

func TestValidate(t *testing.T) {
	err := Validate(10)
	if nil != err {
		t.Errorf("Test Validate has failed : %v", err)
	}

	err = Validate(true)
	if nil != err {
		t.Errorf("Test Validate has failed : %v", err)
	}

	err = Validate(nil)
	if nil == err {
		t.Errorf("Test Validate has failed : %v", err)
	}
}

func TestNewInstance(t *testing.T) {
	channelOption := NewInstance("TestNewInstance", reflect.TypeOf(BOOL_TYPE))
	if !strings.EqualFold(channelOption.Name(), "TestNewInstance") ||
		reflect.Bool != channelOption.t.Kind() {
		t.Errorf("Test NewInstance has failed")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Test NewInstance has sucess")
		}
	}()
	channelOption = NewInstance("TestNewInstance", reflect.TypeOf(BOOL_TYPE))

	t.Errorf("Test NewInstance has failed")
}
