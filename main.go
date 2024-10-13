package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"elif.com/note"
)

//ask user the data that shold be stored
//this data should be written to a file

func main() {
    title, content, err := getNoteData()
    if err != nil {
        fmt.Println(err)
    }
    userNote, err := note.New(title, content)    
    if err != nil{ 
        fmt.Println(err)
    }
    userNote.DisplayContent()
    err = userNote.SaveToFile()
    if err != nil {
        fmt.Println(err)
    }
	
}

func getNoteData () (string, string, error){
    titleInput:= getUserInput("Please enter the title: ")

    contentInput:= getUserInput("Please enter the text you want to save: ")
    return titleInput, contentInput, nil
}

func getUserInput (inputText string) (string){
    fmt.Println(inputText)
    var inputVar string
    
    reader := bufio.NewReader(os.Stdin)

    inputVar, err := reader.ReadString('\n')

    if err != nil {
        return ""
    }
  /*on Windows, a line break is actually not created with just \n,
    but instead with \r\n, so we need to delete both
  */
    inputVar = strings.TrimSuffix(inputVar, "\n")
    inputVar = strings.TrimSuffix(inputVar, "\r")
    return inputVar
}