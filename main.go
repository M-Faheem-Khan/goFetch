package main

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/zcalusic/sysinfo"
)

func main() {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	header, divider := generateHeader(si)

	fmt.Println(header)
	fmt.Println(divider)
	fmt.Printf("OS: %s\n", si.OS.Name)
	fmt.Printf("Kernel: %s\n", si.Kernel.Release)
	fmt.Printf("CPU: %s\n", si.CPU.Model)
	fmt.Printf("Memory: %dmb\n", si.Memory.Size)
}

func generateHeader(si sysinfo.SysInfo) (string, string) {
	user, _ := user.Current()
	header := join(user.Username, "@", si.Node.Hostname)
	divider := generateDivider(len(header))
	return header, divider
}

func generateDivider(l int) string {
	var sb strings.Builder
	for i := 0; i < l; i++ {
		sb.WriteString("-")
	}
	return sb.String()
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
