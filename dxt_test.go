package dxt

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// assertEqual fails if the two values are not equal
func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v != want: %v", got, want)
	}
}

func TestDecodeDXT1(t *testing.T) {
	enc, err := ioutil.ReadFile("testdata/dxt1.encoded")
	assertEqual(t, err == nil, true)

	decg, err := DecodeDXT1(enc, 256, 256)
	assertEqual(t, err == nil, true)

	decw, err := ioutil.ReadFile("testdata/dxt1.decoded")
	assertEqual(t, err == nil, true)

	assertEqual(t, bytes.Compare(decg, decw), 0)
}

func TestDecodeDXT3(t *testing.T) {
	enc, err := ioutil.ReadFile("testdata/dxt3.encoded")
	assertEqual(t, err == nil, true)

	decg, err := DecodeDXT3(enc, 128, 512)
	assertEqual(t, err == nil, true)

	decw, err := ioutil.ReadFile("testdata/dxt3.decoded")
	assertEqual(t, err == nil, true)

	assertEqual(t, bytes.Compare(decg, decw), 0)
}
func TestDecodeDXT5(t *testing.T) {
	enc, err := ioutil.ReadFile("testdata/dxt5.encoded")
	assertEqual(t, err == nil, true)

	decg, err := DecodeDXT5(enc, 64, 64)
	assertEqual(t, err == nil, true)

	decw, err := ioutil.ReadFile("testdata/dxt5.decoded")
	assertEqual(t, err == nil, true)

	assertEqual(t, bytes.Compare(decg, decw), 0)
}

func TestDecodeErrors(t *testing.T) {
	_, err := DecodeDXT1(make([]byte, 0), 4, 4)
	assertEqual(t, err != nil, true)
	_, err = DecodeDXT3(make([]byte, 0), 4, 4)
	assertEqual(t, err != nil, true)
	_, err = DecodeDXT5(make([]byte, 0), 4, 4)
	assertEqual(t, err != nil, true)
}
