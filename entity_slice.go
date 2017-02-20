// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Entity

package ari

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// EntitySlice is a slice of type Entity. Use it where you would use []Entity.
type EntitySlice []Entity

// SortBy returns a new ordered EntitySlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv EntitySlice) SortBy(less func(Entity, Entity) bool) EntitySlice {
	result := make(EntitySlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortEntitySlice(result, less, 0, n, maxDepth)
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapEntitySlice(rcv EntitySlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapEntitySlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapEntitySlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownEntitySlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapEntitySlice(rcv, first, first+i)
		siftDownEntitySlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapEntitySlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapEntitySlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapEntitySlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeEntitySlice(rcv EntitySlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapEntitySlice(rcv, a+i, b+i)
	}
}

func doPivotEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeEntitySlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeEntitySlice(rcv, less, m, m-s, m+s)
		medianOfThreeEntitySlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeEntitySlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapEntitySlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapEntitySlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapEntitySlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeEntitySlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeEntitySlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortEntitySlice(rcv EntitySlice, less func(Entity, Entity) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortEntitySlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotEntitySlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortEntitySlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortEntitySlice(rcv, mhi, b)
		} else {
			quickSortEntitySlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortEntitySlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortEntitySlice(rcv, less, a, b)
	}
}
