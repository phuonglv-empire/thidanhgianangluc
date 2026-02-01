package http

import (
	"github.com/gin-gonic/gin"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/usecase"
	"gorm.io/gorm"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// CORS middleware
	router.Use(corsMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Exam API is running",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Exam routes
		examHandler := NewExamHandler(
			usecase.NewExamUseCase(
				NewExamRepository(db),
			),
		)
		api.GET("/exams", examHandler.GetExams)
		api.GET("/exams/:id", examHandler.GetExam)

		// Exam session routes
		sessionHandler := NewExamSessionHandler(
			usecase.NewExamSessionUseCase(
				NewExamSessionRepository(db),
				NewExamRepository(db),
				NewAnswerRepository(db),
			),
		)
		api.POST("/exam-sessions", sessionHandler.CreateSession)
		api.GET("/exam-sessions/:id", sessionHandler.GetSession)
		api.POST("/exam-sessions/:id/answers", sessionHandler.SubmitAnswer)
		api.POST("/exam-sessions/:id/submit", sessionHandler.SubmitExam)
		api.GET("/exam-sessions/:id/results", sessionHandler.GetResults)
	}
}

// corsMiddleware handles CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
