package main

import (
	"fmt"

	internal "github.com/jordabrell/swap/packages"
)

func main() {
	if internal.FileHomeExist() {
		fmt.Printf("OH! It seems that you do not have the file ~/.aws/credentials\nDo you have awscli installed?\n")
	}
	internal.ReadFile()
}