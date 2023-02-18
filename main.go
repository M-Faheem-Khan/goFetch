package main

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/M-Faheem-Khan/goFetch/utils"
	"github.com/zcalusic/sysinfo"
)

func main() {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	header, divider := generateHeader(si)

	fmt.Println(utils.FormatHeaderString(header))
	fmt.Println(utils.FormatBoldString(divider))
	fmt.Printf("%s: %s\n", utils.FormatKeyString("OS"), si.OS.Name)
	fmt.Printf("%s: %s\n", utils.FormatKeyString("Kernel"), si.Kernel.Release)
	fmt.Printf("%s: %s\n", utils.FormatKeyString("CPU"), si.CPU.Model)
	fmt.Printf("%s: %d mb\n", utils.FormatKeyString("Memory"), si.Memory.Size)
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

// EOF - Changes should reflect
