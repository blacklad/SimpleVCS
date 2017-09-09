package lib

import (
	"bytes"
	"compress/gzip"
	"testing"
)

func TestZip(t *testing.T) {
	data := "testing"
	byteData := []byte(data)
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write(byteData)
	comp.Close()
	zippedData := compBytes.String()
	testData := Zip(byteData)
	if testData != zippedData {
		t.Fail()
	}
}

func TestUnzip(t *testing.T) {
	data := "testing"
	byteData := []byte(data)
	zippedStringData := Zip(byteData)
	zippedData := []byte(zippedStringData)
	var compBytes bytes.Buffer
	compBytes.Write(zippedData)
	comp, _ := gzip.NewReader(&compBytes)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(comp)
	comp.Close()
	if outputBytes.String() != Unzip(zippedData) {
		t.Fail()
	}
}
