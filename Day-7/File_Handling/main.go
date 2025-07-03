package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		//file creation
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file.", err)
			return
		}

		defer file.Close()

		//insert the content
		content := "hello world"
		byte, error := io.WriteString(file, content+"\n")
		fmt.Println("byte written:", byte)
		if error != nil {
			fmt.Println("Error while writing file: ", err)
		}

		fmt.Println("File successfully created....")
	*/

	/*
			//reading file
			file, err := os.Open("example.txt")
			if err != nil {
				fmt.Println("Error while creating file: ", err)
			}
			defer file.Close()

			//create a buffer to read the file content
			buffer := make([]byte, 1024)

			//read the file content into the buffer
			for {
				n, err := file.Read(buffer)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("Error while reading file: ", err)
					return
				}

				//Process the read content
				fmt.Println(string(buffer[:n]))
		}
	*/

	//Read the entire file into a byte slice(where you want to the entire content of a file into memory)
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file.", err)
	}

	fmt.Println(string(content))
}
