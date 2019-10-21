package example_test

import (
	"io/ioutil"
	"testing"

	"github.com/hanpama/mergeb64/example"
)

func Test(t *testing.T) {
	var (
		b   []byte
		err error
	)

	b, err = ioutil.ReadFile("../go.mod")
	if err != nil {
		t.Fatal(err)
		return
	}
	if string(b) != string(example.GoMod) {
		t.Fail()
		return
	}

	b, err = ioutil.ReadFile("../go.sum")
	if err != nil {
		t.Fatal(err)
		return
	}
	if string(b) != string(example.GoSum) {
		t.Fail()
		return
	}
}
