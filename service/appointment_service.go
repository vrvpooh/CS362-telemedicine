package service

import "cs362-telemedicine/model"

type AppointmentService interface {
	CreateAppointment(req model.CreateAppointmentRequest) error
	GetZoomToken(appointmentID string) (string, error)
	UpdateAppointmentStatus(appointmentID string, status string) error
}