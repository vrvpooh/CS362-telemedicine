package repository

import "cs362-telemedicine/model"

type AppointmentRepository interface {
	SaveAppointment(appt *model.Appointment) error
	FindAppointmentByID(id string) (*model.Appointment, error)
	UpdateStatus(id string, status string) error
}