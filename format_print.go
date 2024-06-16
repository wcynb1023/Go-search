package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var srr string

func read() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║   search key:                                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("\033[2C \033[2A \033[11C ")
	input, _ := reader.ReadString('\n')

	str := strings.Fields(input)

	fmt.Print("\033[2B")
	return str
}

func return_web(ans []int) {
	var url_list []string

	for num, i := range ans {
		//fmt.Printf("%d ",i)
		ids := strconv.Itoa(i)
		fmt.Printf("\x1b[34m%3d\x1b[0m : ", num)
		a := find_key(ids)
		fmt.Printf("\033[32m%s\033[0m\n", find_key("title"+ids))
		url_list = append(url_list, a)
	}

	var inp int = 0
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║   go to web:                                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("\033[2C \033[2A \033[11C ")
	fmt.Scanln(&inp)
	cmd := exec.Command("firefox", url_list[inp])

	// 执行命令并捕获输出和错误
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("命令执行出错:", err)
		return
	}
}
