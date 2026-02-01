package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/internal/domain"
	"github.com/phuonglv-empire/thidanhgianangluc/backend/pkg/database"
	"gorm.io/datatypes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Database configuration
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "examuser"),
		Password: getEnv("DB_PASSWORD", "exampass"),
		DBName:   getEnv("DB_NAME", "examdb"),
	}

	// Connect to database
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Starting database seeding...")

	// Create sample exam
	exam := &domain.Exam{
		Title:       "HSA Mock Test 1",
		Description: "Đề thi thử HSA lần 1 - Đánh giá năng lực tư duy",
		Duration:    90, // 90 minutes
		ContentHTML: "<h2>Hướng dẫn làm bài</h2><p>Thời gian làm bài: 90 phút</p><p>Số câu hỏi: 5 câu</p>",
		IsActive:    true,
	}

	if err := db.Create(exam).Error; err != nil {
		log.Fatalf("Failed to create exam: %v", err)
	}
	log.Printf("Created exam: %s (ID: %d)", exam.Title, exam.ID)

	// Create sample questions
	questions := []domain.Question{
		{
			ExamID:      exam.ID,
			ContentHTML: "<p>Cho hàm số f(x) = x² + 2x + 1. Tìm đạo hàm của hàm số?</p>",
			Type:        domain.MultipleChoice,
			Options:     createJSONB([]string{"f'(x) = 2x + 2", "f'(x) = x + 2", "f'(x) = 2x + 1", "f'(x) = x² + 2"}),
			Answer:      createJSONB("f'(x) = 2x + 2"),
			Explanation: "Đạo hàm của x² là 2x, đạo hàm của 2x là 2, đạo hàm của hằng số là 0.",
			Points:      1,
			Order:       1,
		},
		{
			ExamID:      exam.ID,
			ContentHTML: "<p>Trong các câu sau, câu nào đúng về cấu trúc nguyên tử?</p>",
			Type:        domain.MultipleChoice,
			Options:     createJSONB([]string{"Nguyên tử gồm hạt nhân và electron", "Nguyên tử chỉ gồm proton", "Nguyên tử chỉ gồm neutron", "Nguyên tử không có điện tích"}),
			Answer:      createJSONB("Nguyên tử gồm hạt nhân và electron"),
			Explanation: "Nguyên tử gồm hạt nhân (proton và neutron) và các electron quay quanh hạt nhân.",
			Points:      1,
			Order:       2,
		},
		{
			ExamID:      exam.ID,
			ContentHTML: "<p>Tìm giá trị nhỏ nhất của hàm số y = x² - 4x + 5 trên đoạn [0, 3]?</p>",
			Type:        domain.MultipleChoice,
			Options:     createJSONB([]string{"y = 1", "y = 2", "y = 3", "y = 5"}),
			Answer:      createJSONB("y = 1"),
			Explanation: "Hàm số đạt giá trị nhỏ nhất tại x = 2, y(2) = 4 - 8 + 5 = 1",
			Points:      1,
			Order:       3,
		},
		{
			ExamID:      exam.ID,
			ContentHTML: "<p>Phát biểu nào sau đây đúng về quá trình quang hợp?</p>",
			Type:        domain.MultipleChoice,
			Options:     createJSONB([]string{"Xảy ra ở lá cây, tạo ra glucose", "Xảy ra ở rễ cây", "Không cần ánh sáng", "Chỉ xảy ra ban đêm"}),
			Answer:      createJSONB("Xảy ra ở lá cây, tạo ra glucose"),
			Explanation: "Quang hợp xảy ra ở lá cây nhờ chất diệp lục, sử dụng ánh sáng để tạo glucose từ CO2 và H2O.",
			Points:      1,
			Order:       4,
		},
		{
			ExamID:      exam.ID,
			ContentHTML: "<p>Chọn đáp án đúng: Thủ đô của Việt Nam là gì?</p>",
			Type:        domain.MultipleChoice,
			Options:     createJSONB([]string{"Hà Nội", "Hồ Chí Minh", "Đà Nẵng", "Huế"}),
			Answer:      createJSONB("Hà Nội"),
			Explanation: "Hà Nội là thủ đô của nước Cộng hòa Xã hội Chủ nghĩa Việt Nam.",
			Points:      1,
			Order:       5,
		},
	}

	for _, question := range questions {
		if err := db.Create(&question).Error; err != nil {
			log.Printf("Failed to create question: %v", err)
		} else {
			log.Printf("Created question #%d: %s", question.Order, question.ContentHTML[:50])
		}
	}

	log.Println("Database seeding completed successfully!")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func createJSONB(data interface{}) datatypes.JSON {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return nil
	}
	return jsonData
}
