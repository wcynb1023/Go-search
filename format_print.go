package main

import (
	"fmt"
	"strconv"
)

var srr string

/*
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
*/
func return_web(ans []int) []Document {
	url_list := make([]Document, len(ans))
	for num, i := range ans {
		ids := strconv.Itoa(i)
		fmt.Printf("\x1b[34m%3d\x1b[0m : \033[32m%s\033[0m\n", num, find_key("title"+ids))
		url_list[num].URL = find_key(ids)
		url_list[num].Title = find_key("title" + ids)
		url_list[num].ID = ids
	}
	/*
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Println("║   go to web:                                ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("\033[2C \033[2A \033[11C ")
		var inp int
		fmt.Scan(&inp)
		cmd := exec.Command("firefox", url_list[inp].URL)

		// Start the command and ignore stdout and stderr
		if err := cmd.Start(); err != nil {
			fmt.Println("命令执行出错:", err)
			return nil
		}
		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			fmt.Println("命令执行出错:", err)
			return nil
		}
	*/
	return url_list
}
