package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Counter struct {
	wordCount,lineCount,punctuationCount,vowelCount int
}

func main() {
	
		fileContent,err:= fileReader("file.txt")
		if err != nil {
			fmt.Print("Error : ",err)
		}
	startTime := time.Now()

	
	
	 if len(os.Args)<1 {
		fmt.Print("Please enter Num of Go routines you want to create")
	 }
	goRoutines ,err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Print("Invalid Go routine number entered")
	}

	channel := make(chan Counter)

	chunk := len(fileContent) / goRoutines

	fmt.Printf("File Chunks = %v",goRoutines)

	fmt.Printf("/n")

	for i := 0; i < goRoutines; i++ {
		start := i * chunk
		end := (i + 1) * chunk
		go reader(fileContent[start:end], channel)
		counts := <-channel
		fmt.Printf("No of Words of Chunk %d: %d \n", i+1, counts.wordCount)
		fmt.Printf("No of Lines of Chunk %d: %d \n", i+1, counts.lineCount)
		fmt.Printf("No of Vowels of Chunk %d: %d \n", i+1, counts.vowelCount)
		fmt.Printf("No of Punctuation of Chunk %d: %d \n", i+1, counts.punctuationCount)
		fmt.Printf("\n\n")
	}

	fmt.Printf("Execution time: %v\n", time.Since(startTime))


}

func fileReader (filepath string) (string,error)  {
	fileContent ,err := os.ReadFile(filepath)
	if err!=nil {
		return "", err
	}
		Content := string(fileContent)
		return Content, nil
	}


func reader (fileContent string, channel chan Counter ){
	count := Counter{}
for _, char := range fileContent {
    switch {
    case char == ' ' || char == '\t' || char == '\r' || char == '.' || char == ',' || char == ';' || char == ':' || char == '!' || char == '?' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
        count.wordCount++
    case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u':
        count.vowelCount++
    case char == '.' || char == '!' || char == '?' || char == ',' || char == ':' || char == ';' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
        count.punctuationCount++
    case char == '\n':
        count.lineCount++
    }
}
channel<-count
}
