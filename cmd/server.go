package cmd

//go:generate swag init -g server.go

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"    
    "github.com/spf13/cobra"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "main/cmd/docs"
)

// @title MyCLI API
// @description This is a sample server for MyCLI.


// @Summary Root endpoint
// @Description Responds with "Hello, World!"
// @Tags root
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func Hello(c *gin.Context) {
    c.String(200, "Hello, World!")
}

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Run an HTTP server",
    Long:  `This command starts an HTTP server on a specified port.`,
    Run: func(cmd *cobra.Command, args []string) {
        port, _ := cmd.Flags().GetString("port")

        // Set Swagger info dynamically
        docs.SwaggerInfo.Host = ":" + port
        docs.SwaggerInfo.Version = version
                

        r := gin.Default()
		r.SetTrustedProxies(nil)

        r.GET("/", Hello)

        // Serve Swagger UI
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

        r.GET("/swagger", func(c *gin.Context) {
            c.Redirect(302, "/swagger/index.html")
        })

        fmt.Printf("Starting server on port %s...\n", port)
        log.Fatal(r.Run(":" + port))
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
    serverCmd.Flags().StringP("port", "p", "8080", "Port to run the HTTP server on")
}