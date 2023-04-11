package test

import (
	"log"
	"testing"

	"github.com/bobbyirawan/bobbygogo"
	"github.com/go-playground/assert/v2"
)

func TestSort_QuickSort(t *testing.T) {
	arr := []int32{1, 2, 4, 1, 3, 45, 1, 35}
	a := bobbygogo.QuickSort(arr, "asc")
	equal := []int32{1, 1, 1, 2, 3, 4, 35, 45}
	log.Println(a)
	assert.IsEqual(a, equal)
}
