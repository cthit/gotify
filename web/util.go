package web

import "fmt"

func (c *Context) printError(err error) {
	if c.Debug {
		fmt.Println(err)
	} else {
		fmt.Println("An error occurred, turn on debug mode for more info")
	}
}
