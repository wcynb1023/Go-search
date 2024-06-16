package main

import (
	"fmt"
)

func main() {
	if err := initDB(); err != nil {
		fmt.Println("init error")
		return
	}

	c := init_crawl()
	err := c.Visit("https://wiki.osdev.org")
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")

	err2 := c.Visit("https://www.nesdev.org/wiki/")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("finish")

	str := read()

	//ans := merge_list(str)
	ans := merge_list_dc(str, 0, len(str)-1)
	return_web(ans)

}
