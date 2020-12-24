package encoding_test

import (
	"testing"

	"github.com/goframeworks/encoding"
)

func TestEncodingFactory(t *testing.T) {
	if encoding.NewFactory().JsonEncoding() == nil {
		t.Fatal("new json encoding failed, it cannot be nil")
	}
}
