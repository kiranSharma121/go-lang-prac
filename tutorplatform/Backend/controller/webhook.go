package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/webhook"
	"github.com/tutorplatform/database"
	"github.com/tutorplatform/model"
)

func Webhook(c *gin.Context) {
	const MaxBodyBytes = int64(65535)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "read error",
		})
		return
	}
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid signature",
			"error":   err.Error(),
		})
		return
	}
	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid session",
				"error":   err.Error(),
			})
			return
		}
		courseIDStr := session.Metadata["course_id"]
		studentIDstr := session.Metadata["student_id"]
		courseID, _ := strconv.Atoi(courseIDStr)
		studentID, _ := strconv.Atoi(studentIDstr)
		enrollment := model.Enrollment{
			CourseID:  uint(courseID),
			StudentID: uint(studentID),
		}
		err = database.DB.Create(&enrollment).Error
		if err != nil {
			log.Println("enrollment error", err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "enroll success"})

}
