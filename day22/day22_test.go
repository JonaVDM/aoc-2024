package day22_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day22"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`1
10
100
2024`, "\n")
	assert.Equal(t, [2]interface{}{37327623, 0}, day22.Run(file))
}

func TestGetScore(t *testing.T) {
	testCases := []struct {
		start, amount, out int
	}{
		{123, 1, 15887950},
		{123, 2, 16495136},
		{123, 3, 527345},
		{1, 2000, 8685429},
		{10, 2000, 4700978},
		{100, 2000, 15273692},
		{2024, 2000, 8667524},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.start, tc.amount), func(t *testing.T) {
			secret, _ := day22.GetScore(tc.start, tc.amount)
			assert.Equal(t, tc.out, secret)
		})
	}
}

func TestMix(t *testing.T) {
	testCases := []struct {
		tmp, secret, out int
	}{
		{42, 15, 37},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.tmp, tc.secret), func(t *testing.T) {
			assert.Equal(t, tc.out, day22.Mix(tc.tmp, tc.secret))
		})
	}
}

func TestPrune(t *testing.T) {
	testCases := []struct {
		secret, out int
	}{
		{100000000, 16113920},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.secret), func(t *testing.T) {
			assert.Equal(t, tc.out, day22.Prune(tc.secret))
		})
	}
}

func TestGetLastDigit(t *testing.T) {
	testCases := []struct {
		num, out int
	}{
		{12312312310, 0},
		{879131927, 7},
		{123, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.num), func(t *testing.T) {
			assert.Equal(t, tc.out, day22.GetLastDigit(tc.num))
		})
	}
}
