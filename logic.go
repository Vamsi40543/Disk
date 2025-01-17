package main

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

func main() {
	fmt.Println("Disk Usage Information:")
	if runtime.GOOS == "windows" {
		getDiskUsageWindows()
	}
}

func getDiskUsageWindows() {
	drives := []string{"C:\\", "D:\\", "E:\\"}

	for _, drive := range drives {
		h := syscall.MustLoadDLL("kernel32.dll")
		c := h.MustFindProc("GetDiskFreeSpaceExW")

		var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes int64
		_, _, err := c.Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drive))),
			uintptr(unsafe.Pointer(&freeBytesAvailable)),
			uintptr(unsafe.Pointer(&totalNumberOfBytes)),
			uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
		)
		if err.Error() != "The operation completed successfully." {
			fmt.Printf("Error getting stats for %s: %v\n", drive, err)
			continue
		}

		fmt.Printf("Drive: %s\n", drive)
		fmt.Printf("  Total: %.2f GB\n", float64(totalNumberOfBytes)/1e9)
		fmt.Printf("  Used: %.2f GB\n", float64(totalNumberOfBytes-freeBytesAvailable)/1e9)
		fmt.Printf("  Available: %.2f GB\n", float64(freeBytesAvailable)/1e9)
		fmt.Println()
	}
}
