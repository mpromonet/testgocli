package cmd

//go:generate swag init -g server.go

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"main/cmd/docs"
)

// @title MyCLI API
// @description This is a sample server for MyCLI.

// @Summary Root endpoint
// @Description Responds with "Hello, World!"
// @Tags root
// @Produce  json
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func Hello(c *gin.Context) {
	c.JSON(200, "Hello, World!")
}

// @Summary Version endpoint
// @Description Responds with the current version of the application
// @Tags version
// @Produce text/plain
// @Success 200 {string} string "version"
// @Router /version [get]
func Version(c *gin.Context) {
	c.String(200, version)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")

		// Set Swagger info dynamically
		docs.SwaggerInfo.Version = version

		r := gin.Default()
		r.SetTrustedProxies(nil)
		r.Use(cors.Default())

		r.GET("/", Hello)
		r.GET("/version", Version)

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
