package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

var count int = 0
var line int = 0
var test int = 0
var text string = "NULL"

func main() {
	//open the file
	readfile, err := os.Open("markers.txt")

	if err != nil {
		fmt.Println(err)
	}

	arg := os.Args[1]
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	//print line by line
	for test == 0 {
		for fileScanner.Scan() {
			count++
			fmt.Println(count, fileScanner.Text())
			test++
		}
	}

	//reset and open again the file
	readfile, err = os.Open("markers.txt")

	if err != nil {
		fmt.Println(err)
	}

	arg = os.Args[1]

	fileScanner = bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	if arg == "go" {
		fmt.Printf("you chose: ")
		fmt.Scanln(&line)
		count = 0
		for fileScanner.Scan() {
			count++
			if count == line {
				link := fileScanner.Text()
				fmt.Println(count, fileScanner.Text())
				openbrowser(link)
			}
		}

	}

	if arg == "append" {
		readfile, err = os.OpenFile("markers.txt", os.O_APPEND, 0644)
		fmt.Printf("Write a full link ")
		fmt.Scanln(&text)

		_, err2 := readfile.WriteString(text + "\n")
		if err2 != nil {
			fmt.Println("Could not append the text")
		} else {
			fmt.Println("Operation successful! Text has been appended to example.txt")
		}
		readfile.Close()
	}
}
