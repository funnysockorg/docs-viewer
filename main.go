package main

import (
	"net/http"
    "fmt"
    "os"
    "bufio"
    "strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()
    }
}

func main() {
    router := gin.Default()
    router.Use(Auth())

    router.LoadHTMLFiles("auth/index.html")

    router.GET("/auth", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "index.html", gin.H{})
    })

    router.POST("/auth", func(ctx *gin.Context) {
        fmt.Println(getAuthDataFromEnv())
    })

    router.Run(":8080")
}

func getAuthDataFromEnv() map[string]string {
    envData := map[string]string{}

    file, err := os.Open(".env")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        readData := strings.Split(scanner.Text(), "=")

        envData[readData[0]] = readData[1]
    }
    
    return envData
}

