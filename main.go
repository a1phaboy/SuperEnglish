package main

import (
	"fmt"
	"os"
)

func main() {
	var option int
	showBanner()
	for{
		showBash()
		fmt.Scanf("%d",&option)
		switch option {
		case 1:
			addWord()
			break
		case 2:
			test()
			break
		case 3:
			quit()
		}

	}
}
func quit(){
	os.Exit(0)
}
