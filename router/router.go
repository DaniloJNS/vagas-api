package router

import "github.com/gin-gonic/gin"

func Initializer() {
  // Initializer Router
  r := gin.Default()
  // define an routes to GET /ping

  // Setup routes
  InitializeRouters(r)

  // Initialize WEB SERVER
  r.Run(":4000")
}
