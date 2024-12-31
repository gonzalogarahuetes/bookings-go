package dbrepo

import (
	"errors"
	"slices"
	"time"

	"github.com/gonzalogarahuetes/bookings-go/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// Inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID != 1 {
		return 0, errors.New("reservation could not be inserted")
	}
	return 1, nil
}

// InsertRoomRestriction Inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("restriction could not be inserted")
	}
	return nil
}

// SearchAvailabilityForDatesByRoomID Returns true if availability exists for a certain room
func (m *testDBRepo) SearchAvailabilityForDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	allowedRoomIDs := make([]int, 1, 2)
	if !slices.Contains(allowedRoomIDs, roomID) {
		return false, errors.New("roomID is not valid")
	}
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms if any for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	startDate, _ := time.Parse("2006-01-02", "2033-02-02")
	if start == startDate {
		return rooms, errors.New("start date is not valid")
	}
	rooms = append(rooms, models.Room{
		ID:       1,
		RoomName: "General's Quarters",
	})
	return rooms, nil
}

// Gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}
