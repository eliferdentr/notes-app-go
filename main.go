package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"elif.com/note"
    "elif.com/todo"
)


type Saver interface {
    SaveToFile() error
}

type outputtapbe interface {
    Saver
    DisplayContent()
}

//ask user the data that shold be stored
//this data should be written to a file

func main() {

    //try note
    title, content, err := getNoteData()
    if err != nil {
        fmt.Println(err)
        return
    }
    userNote, err := note.New(title, content)    
    if err != nil{ 
        fmt.Println(err)
        return
    }
    
    err = outputData(userNote)
    if err != nil {
        fmt.Println(err)
        return
    }

    //try todo
    todoText := getTodoData()
    todo, err := todo.New(todoText)

    if err != nil {
        fmt.Println(err)
        return
    }

    err = outputData(todo)

    if err != nil {
        fmt.Println(err)
        return
    }
	
}

func getNoteData () (string, string, error){
    titleInput:= getUserInput("Please enter the title: ")

    contentInput:= getUserInput("Please enter the text you want to save: ")
    return titleInput, contentInput, nil
}

func getTodoData() (string){
    text := getUserInput("Enter TODO: ")
    return text
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

func outputData(data outputtapbe) error{
    data.DisplayContent()
    err := saveData(data)
    if err != nil {
        return err
    }
    return nil
}

func saveData (data Saver) error{
    err := data.SaveToFile()
    if err != nil {
        fmt.Println("Saving the data has failed.")
        return err
    }

    fmt.Println("File has been successfully saved")
    return nil
}

func getTypeOfTheValue (value interface{}){
    if typedVal, ok := value.(int) ; ok {
        fmt.Println("Integer: ", typedVal)
        return
    }

    if typedVal, ok := value.(float64) ; ok {
        fmt.Println("Float: ", typedVal)
        return
    }
}