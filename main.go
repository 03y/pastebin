package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

const VERSION           string = "0.0.1-Alpha"
const PASTE_DESTINATION string = "/var/www/html/paste/"
const LISTEN            string = "localhost:6600"

type paste struct {
    ID      string  `json:"id"`
    Text    string  `json:"text"`
}

func pasteIt(c *gin.Context) {
    var newPaste paste

    if err := c.BindJSON(&newPaste); err != nil {
        return
    }

    // Log
    fmt.Printf("New paste (ID: %s)\n", newPaste.ID)

    data := []byte(newPaste.Text)
    err := os.WriteFile(PASTE_DESTINATION + newPaste.ID, data, 0644)
    if err != nil {
        fmt.Println("Failed to write file:")
        fmt.Println(err)
    }

    // Respond
    c.IndentedJSON(http.StatusCreated, newPaste)
}

func main() {
    fmt.Println("Paste backend " + VERSION) 

    router := gin.Default()
    router.POST("/new", pasteIt)
    fmt.Printf("\nListening on %s...\n", LISTEN)
    router.Run(LISTEN)
}

