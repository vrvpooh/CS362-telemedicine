package handler

import "github.com/gin-gonic/gin"

type AppointmentHandler interface {
	CreateAppointment(c *gin.Context)
	GetZoomToken(c *gin.Context)
	UpdateAppointmentStatus(c *gin.Context)
}