// Code generated by 'go generate'; DO NOT EDIT.

package dos

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modole32    = windows.NewLazySystemDLL("ole32.dll")

	procCopyFileW           = modkernel32.NewProc("CopyFileW")
	procGetDiskFreeSpaceExW = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procCoInitializeEx      = modole32.NewProc("CoInitializeEx")
	procCoUninitialize      = modole32.NewProc("CoUninitialize")
)

func CopyFile(src string, dst string, isFailIfExist bool) (n uint32, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(src)
	if err != nil {
		return
	}
	var _p1 *uint16
	_p1, err = syscall.UTF16PtrFromString(dst)
	if err != nil {
		return
	}
	return _CopyFile(_p0, _p1, isFailIfExist)
}

func _CopyFile(src *uint16, dst *uint16, isFailIfExist bool) (n uint32, err error) {
	var _p2 uint32
	if isFailIfExist {
		_p2 = 1
	} else {
		_p2 = 0
	}
	r0, _, e1 := syscall.Syscall(procCopyFileW.Addr(), 3, uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dst)), uintptr(_p2))
	n = uint32(r0)
	if n == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getDiskFreeSpaceEx(rootPathName string, free *uint64, total *uint64, totalFree *uint64) (n uint32, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(rootPathName)
	if err != nil {
		return
	}
	return _getDiskFreeSpaceEx(_p0, free, total, totalFree)
}

func _getDiskFreeSpaceEx(rootPathName *uint16, free *uint64, total *uint64, totalFree *uint64) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetDiskFreeSpaceExW.Addr(), 4, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(free)), uintptr(unsafe.Pointer(total)), uintptr(unsafe.Pointer(totalFree)), 0, 0)
	n = uint32(r0)
	if n == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CoInitializeEx(res uintptr, opt uintptr) {
	syscall.Syscall(procCoInitializeEx.Addr(), 2, uintptr(res), uintptr(opt), 0)
	return
}

func CoUninitialize() {
	syscall.Syscall(procCoUninitialize.Addr(), 0, 0, 0, 0)
	return
}
