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

	api := s.Router.Group("/api/queue")

	api.Use(s.GenerateRequestID)
	api.Use(s.Authenticate)
	api.POST("/join", s.qh.JoinQueue)
	api.GET("/upcoming", s.qh.GetCurrentQueueNumber)
	api.GET("/next", s.qh.RetrieveNextInQueue)

}

func (s *Server) GenerateRequestID(c *gin.Context) {
	log.Println("Generating request ID")
	requestID := uuid.New().String()
	log.Println("Generated request ID:", requestID)
	c.Set("request_id", requestID)
	c.Header("X-Request-ID", requestID)
	c.Next()
}

func (s *Server) Authenticate(c *gin.Context) {
	log.Println("authenticating user...")

	req, err := http.NewRequestWithContext(c, "GET", "http://user-ms:3005/api/users/validatejwt", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "try again later.")
		c.Abort()
		return
	}

	// extract the "Bearer XXXX" and set as header
	token := c.GetHeader("Authorisation")
	if token == "" {
		c.String(http.StatusBadRequest, "missing auth token please login.")
		c.Abort()
		return
	}

	req.Header.Set("Authorisation", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "try again later.")
		c.Abort()
		return
	}

	defer resp.Body.Close()

	log.Printf("res %+v", resp)
	if resp.Status == http.StatusText(http.StatusUnauthorized) {
		c.String(http.StatusUnauthorized, "Unauthorised. Go be authorised den come back")
		c.Abort()
		return
	}
	c.Next()
}
