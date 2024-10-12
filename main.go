package main

import (
	"errors"
	"fmt"
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
	
}

func getNoteData () (string, string, error){
    titleInput, err := getUserInput("Please enter the title: ")
    if err != nil{ 
        return "", "", err
    }

    contentInput, err := getUserInput("Please enter the text you want to save: ")
    if err != nil{ 
        return "", "", err
    }

    return titleInput, contentInput, nil
}

func getUserInput (inputText string) (string, error){
    fmt.Println(inputText)
    var inputVar string
    fmt.Scan(&inputVar)
    if inputText == "" {
        return "", errors.New("You cannot enter empty string")
    }
    return inputVar, nil
}