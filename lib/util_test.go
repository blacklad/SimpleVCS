package lib

import (
	"bytes"
	"compress/gzip"
	"testing"
)

const data = "testing"

func TestZip(t *testing.T) {
	byteData := []byte(data)
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write(byteData)
	comp.Close()
	zippedData := compBytes.String()
	testData := Zip(data)
	if testData != zippedData {
		t.Fail()
	}
}

func TestUnzip(t *testing.T) {
	zippedStringData := Zip(data)
	zippedData := []byte(zippedStringData)
	var compBytes bytes.Buffer
	compBytes.Write(zippedData)
	comp, _ := gzip.NewReader(&compBytes)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(comp)
	comp.Close()
	if outputBytes.String() != Unzip(string(zippedData)) {
		t.Fail()
	}
}
