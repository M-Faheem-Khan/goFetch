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

	header = utils.FormatHeaderString(header)
	divider = utils.FormatBoldString(divider)
	os := fmt.Sprintf("%s: %s", utils.FormatKeyString("OS"), si.OS.Name)
	kernel := fmt.Sprintf("%s: %s", utils.FormatKeyString("Kernel"), si.Kernel.Release)
	cpu := fmt.Sprintf("%s: %s", utils.FormatKeyString("CPU"), si.CPU.Model)
	memory := fmt.Sprintf("%s: %d mb", utils.FormatKeyString("Memory"), si.Memory.Size)

	fmt.Println(utils.GenerateOutput(header, divider, os, kernel, cpu, memory))
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
