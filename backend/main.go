package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    cors "github.com/rs/cors/wrapper/gin"
)

const VERSION           string = "0.0.1-Alpha"
const PASTE_DESTINATION string = "/var/www/html/paste/"

type paste struct {
    ID      string  `json:"id"`
    Text    string  `json:"text"`
}

func pasteIt(c *gin.Context) {
    var newPaste paste

    if err := c.BindJSON(&newPaste); err != nil {
        return
    }

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
    fmt.Printf("Paste backend %s\n", VERSION) 
   
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./paste_backend <URL>:<PORT>")
        return
    }

    var listen string = os.Args[1]

    router := gin.Default()
    router.Use(cors.AllowAll())
    router.POST("/new", pasteIt)
    fmt.Printf("\nListening on %s...\n", listen)
    router.Run(listen)
}

