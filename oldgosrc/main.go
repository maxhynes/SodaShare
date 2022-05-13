package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg = sync.WaitGroup{}

func main() {
	for {

		fmt.Println("Invoking webserver into different thread...")
		go databaseInteraction("this is a test")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		if strings.TrimRight(text, "\n") == "quit" { // Applicaiton Exit
			fmt.Println("Exiting Application")
			os.Exit(1)
		} else {
			fmt.Println(text)
		}
	}
}
func databaseInteraction(input string) {

	fmt.Println(input)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		c.FileAttachment("file.zip", "testfile.zip")

	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
