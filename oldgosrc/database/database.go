package main

import (
	"bufio"
	"fmt"
	"os"
)

func Database(SCHEMA string) {

	fmt.Println(SCHEMA)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	Database(text)

}
