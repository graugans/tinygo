// +build gc.none

package runtime

// This GC strategy provides no memory allocation at all. It can be useful to
// detect where in a program memory is allocated, or to combine this runtime
// with a separate (external) garbage collector.

import (
	"unsafe"
)

func alloc(size uintptr) unsafe.Pointer

func free(ptr unsafe.Pointer) {
	// Nothing to free when nothing gets allocated.
}

func GC() {
	// Unimplemented.
}
