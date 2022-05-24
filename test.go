package main

import "fmt"

func main() {
	cmd := "ls"
	args := []string{"-l"}
	//cmd := exec.Command("ls", "/home/zt", "-l")
	//cmd := exec.Command("uname -r")
	file := "/home/zt/code/go/src/multiTaskCommand/1"
	fileline := nslookup(file, 16)
	for i := 0; i < len(fileline); i++ {
		for j := 0; j < len(fileline[i]); j++ {
			fmt.Println(fileline[i][j])
		}
	}
	fmt.Println()
	linuxCommandGoroutine(cmd, args, 16, fileline)
}
