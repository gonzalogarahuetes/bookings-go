package repository

import "github.com/gonzalogarahuetes/bookings-go/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
