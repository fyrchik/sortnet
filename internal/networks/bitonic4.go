// Code generated by "sortnet"; DO NOT EDIT.
package networks

import "sort"

// Bitonic4 sorts 4-element slice in an oblivious fashion.
func Bitonic4(a sort.Interface) {
	if !a.Less(0, 1) {
		a.Swap(0, 1)
	}
	if !a.Less(2, 3) {
		a.Swap(2, 3)
	}
	if !a.Less(0, 3) {
		a.Swap(0, 3)
	}
	if !a.Less(1, 2) {
		a.Swap(1, 2)
	}
	if !a.Less(0, 1) {
		a.Swap(0, 1)
	}
	if !a.Less(2, 3) {
		a.Swap(2, 3)
	}
}
