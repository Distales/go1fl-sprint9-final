package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 1000} {
		data := generateRandomElements(size)
		want := size
		if want < 0 {
			want = 0
		}
		if len(data) != want {
			t.Fatalf("size %d: got length %d", size, len(data))
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("nonpositive value: %d", value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"one", []int{7}, 7},
		{"first", []int{9, 1, 2}, 9}, {"last", []int{1, 2, 9}, 9},
		{"duplicates", []int{5, 5, 5}, 5}, {"negative", []int{-5, -2, -8}, -2},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for name, find := range map[string]func([]int) int{"sequential": maximum, "parallel": maxChunks} {
				if got := find(tc.data); got != tc.want {
					t.Errorf("%s: got %d, want %d", name, got, tc.want)
				}
			}
		})
	}
}

func TestMaxChunksPartitions(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for size := 1; size <= 100; size++ {
		data := make([]int, size)
		for i := range data {
			data[i] = rng.Intn(1000) - 500
		}
		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Fatalf("size %d: got %d, want %d", size, got, want)
		}
		for i := range data {
			old := data[i]
			data[i] = 1001
			if got := maxChunks(data); got != 1001 {
				t.Fatalf("size %d index %d: lost maximum", size, i)
			}
			data[i] = old
		}
	}
}
