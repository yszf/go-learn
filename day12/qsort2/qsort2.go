package qsort2

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"
import (
	"unsafe"
)

// 闭包函数无法导出为C语言函数，因此无法直接将闭包函数传入C语言的qsort函数
func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	go_qsort_compare_info.fn = cmp

	C.qsort(base, C.size_t(num), C.size_t(size),
		// cmp, // cannot use cmp (variable of type func(a unsafe.Pointer, b unsafe.Pointer) int) as *[0]byte value in argument to C.qsort
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}
