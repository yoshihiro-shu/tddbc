package tddbc_test

import (
	"testing"

	"github.com/yoshihiro-shu/tddbc"
)

func TestSayHello(t *testing.T) {
	actual := tddbc.SayHello()
	if actual != "Hello, World!" {
		t.Errorf("got %v\nwant %v", actual, "Hello, World!")
	}
}
