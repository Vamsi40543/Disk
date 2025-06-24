# Disk Usage Information (Windows)

This Go program prints disk usage information for selected Windows drives (`C:\`, `D:\`, and `E:\`). It uses low-level Windows system calls to retrieve total, used, and available space.

##  Features

- Uses Windows API (`GetDiskFreeSpaceExW`) via `syscall`
- Displays total, used, and available disk space in GB
- Supports multiple drives (`C:\`, `D:\`, `E:\` by default)

##  Platform Support

-  Windows only

##  Prerequisites

- Go 1.18 or higher
- Windows OS

##  How to Run

1. Save the code to a file, e.g., `main.go`
2. Open a terminal (Command Prompt or PowerShell)
3. Navigate to the file's directory
4. Run the program:

```bash
go run main.go
```
##Sample Output
Disk Usage Information:
Drive: C:\
  Total: 256.00 GB
  Used: 120.50 GB
  Available: 135.50 GB

Drive: D:\
  Total: 512.00 GB
  Used: 200.75 GB
  Available: 311.25 GB

Drive: E:\
  Error getting stats for E:\: The system cannot find the drive specified.
