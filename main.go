package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"os/exec"
)

var srr string

func main() {

	if err := initClient(); err != nil {
		fmt.Println("init error");
		return
	}
	
	/* c := init_crawl()
	err := c.Visit("https://wiki.osdev.org")
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")  */
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║   search key:                                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("\033[2C \033[2A \033[11C ")
	input, _ := reader.ReadString('\n')
	str := strings.Fields(input)

	fmt.Print("\033[2B")
	var ans []int
		arr := Get_list(str[0])
		intSlice1 := make([]int, len(arr))
		for j , val := range arr{
			intSlice1[j] , _ = strconv.Atoi(val)
		}
		ans = merge(intSlice1,intSlice1)

	for i := 1 ; i < len(str) ; i++{
		arr := Get_list(str[i])
		intSlice1 := make([]int, len(arr))
		for j , val := range arr{
			intSlice1[j] , _ = strconv.Atoi(val)
		}
		ans = merge(ans,intSlice1)
	}

	var url_list []string
	
	for num , i := range ans{
		//fmt.Printf("%d ",i)
		ids := strconv.Itoa(i)
		fmt.Printf("👍👍👍👍\x1b[34m%3d\x1b[0m : ",num)
		a := find_key(ids)
		url_list = append(url_list , a)
	}  
	
	var inp int = 0
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║   go to web:                                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("\033[2C \033[2A \033[11C ")
	fmt.Scanln(&inp)
	cmd := exec.Command("firefox",url_list[inp] )

	// 执行命令并捕获输出和错误
	_ , err := cmd.Output()
	if err != nil {
		fmt.Println("命令执行出错:", err)
		return
	}
	
}
