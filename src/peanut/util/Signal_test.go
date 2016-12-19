package util

import (
	"strings"
	"testing"
)

func TestValueOf_Signal(t *testing.T) {
	signal_1 := ValueOf_Signal("TEST_SIGNAL")
	signal_2 := ValueOf_Signal("TEST_SIGNAL")

	if !strings.EqualFold(signal_1.Name(), "TEST_SIGNAL") ||
		!strings.EqualFold(signal_2.Name(), "TEST_SIGNAL") ||
		signal_1.Id() != signal_2.Id() {
		t.Errorf("Test ValueOf_Signal has failed")
	}
}

func TestExpect(t *testing.T) {
	signal_1 := ValueOf_Signal("TEST_SIGNAL")
	signal_2 := ValueOf_Signal("TEST_SIGNAL")

	if err := signal_1.Expect(signal_2); nil != err {
		t.Errorf("Test Expect has failed")
	}
}

func TestEqual(t *testing.T) {
	signal_1 := ValueOf_Signal("TEST_SIGNAL")
	signal_2 := ValueOf_Signal("TEST_SIGNAL")

	if !signal_1.Equal(signal_2) {
		t.Errorf("Test Equal has failed")
	}

	var integerVal int
	if signal_1.Equal(&integerVal) {
		t.Errorf("Test Equal has failed")
	}
}

// TODO : After compareTo func implementation
func TestCompareTo(t *testing.T) {
	signal_1 := ValueOf_Signal("TEST_SIGNAL")
	signal_2 := ValueOf_Signal("TEST_SIGNAL")

	if !signal_1.Equal(signal_2) {
		t.Errorf("Test Equal has failed")
	}

	var integerVal int
	if signal_1.Equal(&integerVal) {
		t.Errorf("Test Equal has failed")
	}
}
