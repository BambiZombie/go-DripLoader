package main

// 使用golang的syscall包调用api
/*
import (
	"syscall"
	"unsafe"
)

var (
	ntdll                       = syscall.NewLazyDLL(string([]byte{'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'}))
	procNTAllocateVirtualMemory = ntdll.NewProc("NtAllocateVirtualMemory")
	procNTWriteVirtualMemory    = ntdll.NewProc("NtWriteVirtualMemory")
	procNTProtectVirtualMemory  = ntdll.NewProc("NtProtectVirtualMemory")
	procNTCreateThreadEx        = ntdll.NewProc("NtCreateThreadEx")
)

func NTAllocateVirtualMemory(hProcess uintptr, lpAddress *uintptr, zerobits uintptr, dwSize *uint32, flAllocationType uint32, flProtect uint32) {
	_, _, err := syscall.Syscall6(procNTAllocateVirtualMemory.Addr(), 6, uintptr(hProcess), uintptr(unsafe.Pointer(lpAddress)), uintptr(zerobits), uintptr(unsafe.Pointer(dwSize)), uintptr(flAllocationType), uintptr(flProtect))
	print(err)
	print("SYSCALL: NtAllocateVirtualMemory", "hProcess=", hProcess, ", ", "lpAddress=", lpAddress, ", ", "zerobits=", zerobits, ", ", "dwSize=", dwSize, ", ", "flAllocationType=", flAllocationType, ", ", "flProtect=", flProtect, "\n")
	return
}

func NTWriteVirtualMemory(hProcess uintptr, lpBaseAddress uintptr, lpBuffer *byte, nSize uintptr, lpNumberOfBytesWritten *uintptr) {
	_, _, err := syscall.Syscall6(procNTWriteVirtualMemory.Addr(), 5, uintptr(hProcess), uintptr(lpBaseAddress), uintptr(unsafe.Pointer(lpBuffer)), uintptr(nSize), uintptr(unsafe.Pointer(lpNumberOfBytesWritten)), 0)
	print(err)
	print("SYSCALL: NtWriteVirtualMemory", "hProcess=", hProcess, ", ", "lpBaseAddress=", lpBaseAddress, ", ", "lpBuffer=", lpBuffer, ", ", "nSize=", nSize, ", ", "lpNumberOfBytesWritten=", lpNumberOfBytesWritten, "\n")
	return
}

func NTProtectVirtualMemory(hProcess uintptr, lpBaseAddress *uintptr, dwSize *uint32, flProtect uint32, oldProtect *uintptr) {
	_, _, err := syscall.Syscall6(procNTProtectVirtualMemory.Addr(), 5, uintptr(hProcess), uintptr(unsafe.Pointer(lpBaseAddress)), uintptr(unsafe.Pointer(dwSize)), uintptr(flProtect), uintptr(unsafe.Pointer(oldProtect)), 0)
	print(err)
	print("SYSCALL: NtProtectVirtualMemory", "hProcess=", hProcess, ", ", "lpBaseAddress=", lpBaseAddress, ", ", "dwSize=", dwSize, ", ", "flProtect=", flProtect, ", ", "oldProtect=", oldProtect, "\n")
	return
}

func NTCreateThreadEx(hThread *uintptr, desiredaccess uintptr, objattrib uintptr, processhandle uintptr, lpstartaddr uintptr, lpparam uintptr, createsuspended uintptr, zerobits uintptr, sizeofstack uintptr, sizeofstackreserve uintptr, lpbytesbuffer uintptr) {
	_, _, err := syscall.Syscall12(procNTCreateThreadEx.Addr(), 11, uintptr(unsafe.Pointer(hThread)), uintptr(desiredaccess), uintptr(objattrib), uintptr(processhandle), uintptr(lpstartaddr), uintptr(lpparam), uintptr(createsuspended), uintptr(zerobits), uintptr(sizeofstack), uintptr(sizeofstackreserve), uintptr(lpbytesbuffer), 0)
	print(err)
	print("SYSCALL: NtCreateThreadEx(", "hThread=", hThread, ", ", "desiredaccess=", desiredaccess, ", ", "objattrib=", objattrib, ", ", "processhandle=", processhandle, ", ", "lpstartaddr=", lpstartaddr, ", ", "lpparam=", lpparam, ", ", "createsuspended=", createsuspended, ", ", "zerobits=", zerobits, ", ", "sizeofstack=", sizeofstack, ", ", "sizeofstackreserve=", sizeofstackreserve, ", ", "lpbytesbuffer=", lpbytesbuffer, "\n")
	return
}
*/
