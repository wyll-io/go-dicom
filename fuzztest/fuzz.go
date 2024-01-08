package fuzz

import (
	"bytes"

	dicom "github.com/wyll-io/go-dicom"
)

func Fuzz(data []byte) int {
	_, _ = dicom.ReadDataSet(bytes.NewBuffer(data), dicom.ReadOptions{})
	return 1
}
