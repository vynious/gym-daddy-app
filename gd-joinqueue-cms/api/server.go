package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vynious/gd-joinqueue-cms/logger"
)

type Server struct {
	Router *gin.Engine
	qh     *QueueHandler
}

func SpawnServer() *Server {
	engine := gin.Default()
	prod := logger.SpawnKafkaProducer(logger.LoadKafkaConfigurations())
	qh := SpawnQueueHandler(prod)
	return &Server{
		Router: engine,
		qh:     qh,
	}
}

func (s *Server) MountHandlers() {
	
	s.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	s.Router.Use(CORSMiddleware())
	api := s.Router.Group("/api/queue")

	api.Use(s.GenerateRequestID)

	
	api.POST("/join",s.AuthenticateDefault, s.qh.JoinQueue)
	api.GET("/upcoming",s.AuthenticateDefault, s.qh.GetCurrentQueueNumber)
	api.GET("/next", s.AuthenticateAdmin ,s.qh.RetrieveNextInQueue)

}

func (s *Server) GenerateRequestID(c *gin.Context) {
	log.Println("Generating request ID")
	requestID := uuid.New().String()
	log.Println("Generated request ID:", requestID)
	c.Set("request_id", requestID)
	c.Header("X-Request-ID", requestID)
	c.Next()
}

func (s *Server) AuthenticateDefault(c *gin.Context) {
	log.Println("authenticating user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt/default", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"missing auth token please login."})
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorised. Go be authorised den come back"})
		c.Abort()
		return
	}
	c.Next()
}


func (s *Server) AuthenticateAdmin(c *gin.Context) {
	log.Println("authenticating user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt/admin", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"missing auth token please login."})
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "try again later."})
		c.Abort()
		return
	}
	if resp.StatusCode == 401 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "only admin-level access"})
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorised. Go be authorised den come back"})
		c.Abort()
		return
	}
	c.Next()
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorisation")

		// Add this line to allow preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}