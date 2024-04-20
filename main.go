package main

import (
	"fmt"
)

func main() {

	if err := initDB(); err != nil {
		fmt.Println("init error")
		return
	}

	/* c := init_crawl()
	err := c.Visit("https://wiki.osdev.org")
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")  */

	str := read()
	ans := merge_list(str)

	return_web(ans)

}
