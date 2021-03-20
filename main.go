package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main(){

    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Yet Another Calculator")

    for {
        fmt.Print("-> ")

        // Get input line and remove \n
        text,_ := reader.ReadString('\n')
        text = strings.ReplaceAll(text, "\n", "")


        if strings.Compare(text, "exit") == 0 {
            break
        }
    }

}
