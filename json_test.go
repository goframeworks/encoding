package encoding_test

import (
	"github.com/goframeworks/encoding"
	"reflect"
	"testing"
)

func TestJsonEncoding(t *testing.T) {
	type Map struct {
		Key   string
		Value string
	}

	var demo = Map{
		Key: "sd",
		Value: "sdfjk",
	}

	b, err := encoding.NewJsonEncoding().Encode(demo)
	if err != nil {
		t.Fatalf("json encode failed, %v", err)
	}

	var decoded Map
	if err = encoding.NewJsonEncoding().Decode(b, &decoded); err != nil {
		t.Fatalf("json decode failed, %v", err)
	}

	if !reflect.DeepEqual(demo, decoded) {
		t.Fatal("result not equal in json encoding")
	}
}