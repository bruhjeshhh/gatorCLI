package main

import (
	"fmt"

	"github.com/bruhjeshhh/gatorCLI/internal/config"
)

func main() {
	resp, _ := config.Read()
	_ = resp.SetUser("brajesh")
	resp, _ = config.Read()
	fmt.Print(resp)

}
