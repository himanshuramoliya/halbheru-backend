package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the ride-sharing platform
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	Password    string `json:"-" gorm:"not null"` // Hidden in JSON responses
	Phone       string `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Gender      string `json:"gender"`
	
	// Relationships
	CreatedRides []Ride `json:"created_rides,omitempty" gorm:"foreignKey:DriverID"`
	JoinedRides  []Ride `json:"joined_rides,omitempty" gorm:"many2many:ride_passengers;"`
}

// Ride represents a ride in the platform
type Ride struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	DriverID     uint      `json:"driver_id" gorm:"not null"`
	Driver       User      `json:"driver" gorm:"foreignKey:DriverID"`
	
	Origin       string    `json:"origin" gorm:"not null"`
	Destination  string    `json:"destination" gorm:"not null"`
	DepartureTime time.Time `json:"departure_time" gorm:"not null"`
	ArrivalTime  *time.Time `json:"arrival_time"`
	
	AvailableSeats int     `json:"available_seats" gorm:"not null"`
	PricePerSeat   float64 `json:"price_per_seat" gorm:"not null"`
	
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'active'"` // active, completed, cancelled
	
	// Location coordinates (optional)
	OriginLat      *float64 `json:"origin_lat"`
	OriginLng      *float64 `json:"origin_lng"`
	DestinationLat *float64 `json:"destination_lat"`
	DestinationLng *float64 `json:"destination_lng"`
	
	// Relationships
	Passengers []User `json:"passengers,omitempty" gorm:"many2many:ride_passengers;"`
}

// RidePassenger represents the many-to-many relationship between rides and passengers
type RidePassenger struct {
	RideID      uint      `json:"ride_id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"primarykey"`
	JoinedAt    time.Time `json:"joined_at" gorm:"autoCreateTime"`
	SeatsBooked int       `json:"seats_booked" gorm:"default:1"`
	Status      string    `json:"status" gorm:"default:'confirmed'"` // confirmed, cancelled
}
