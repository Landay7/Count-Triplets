package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_сountTriplets1(t *testing.T) {
	arr := []int64{1, 2, 2, 4}
	result := countTriplets(arr, 2)
	assert.Equal(t, result, int64(2))
}

func Test_сountTriplets2(t *testing.T) {
	arr := []int64{1, 3, 9, 9, 27, 81}
	result := countTriplets(arr, 3)
	assert.Equal(t, result, int64(6))
}

func Test_сountTriplets3(t *testing.T) {
	arr := []int64{1, 5, 5, 25, 125}
	result := countTriplets(arr, 5)
	assert.Equal(t, result, int64(4))
}