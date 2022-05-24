package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"io"
	"os/exec"
	"sync"
)

/*读取一个文件，并根据输入的核心数，平均分配任务。以便后续
多协程运行*/
func nslookup(filepath string, process int) [][]string {
	fileline := make([][]string, process)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("fail to open the file %s.", filepath)
	}
	buff := bufio.NewReader(file)
	for i:= 0; true; {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		i = i % process
		fileline[i] = append(fileline[i], string(data))
		i++
	}
	return fileline
}

/*多协程运行命令，command是要执行的系统命令,
commandarg是系统命令的参数,context是系统命令要处理的文本*/
func linuxCommandGoroutine(command string, commandarg []string, processNum int, context [][]string){
	var wg sync.WaitGroup
	for i := 0; i < processNum; i++ {
		wg.Add(1)
		go linuxCommand(command, commandarg, context[i], &wg)
	}
	wg.Wait()
	fmt.Println("over")
}

func linuxCommand(command string, commandarg []string, context []string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	length := len(context)
	fmt.Println(length)
	for i := 0; i < length; i++ {
		newCommandarg := []string{context[i]}
		fmt.Println(context[0])
		fmt.Println(commandarg[0])
		newCommandarg = append(newCommandarg, commandarg...)
		fmt.Println()
		cmd := exec.Command(command, newCommandarg...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		for b, err := out.ReadBytes('\n'); err == nil; {
			fmt.Println(string(b))
			b, err = out.ReadBytes('\n')
		}
	}
	fmt.Println()
}
