package models

import (
	"github.com/djfemz/rave/rave-app/security/otp"
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN       = "ADMIN"
	ORGANIZER   = "ORGANIZER"
	EVENT_STAFF = "EVENT_STAFF"
)

const (
	NOT_STARTED = "NOT_STARTED"
	ONGOING     = "ON_GOING"
	ENDED       = "ENDED"
)

const (
	ACTIVE    = "ACTIVE"
	SUSPENDED = "SUSPENDED"
	IN_ACTIVE = "IN_ACTIVE"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(EventStaff{}).String()] = EventStaff{}
	Entities[reflect.ValueOf(Ticket{}).String()] = Ticket{}
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
	Otp       *otp.OneTimePassword `gorm:"embedded;embeddedPrefix:otp"`
	EventId   uint64
	Events    []*Event
}

type User struct {
	ID       uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AdditionalInformationFields []string

type Ticket struct {
	ID                           uint64 `gorm:"primaryKey"`
	Type                         string
	Name                         string                      `json:"name"`
	Stock                        uint64                      `json:"stock"`
	NumberAvailable              uint64                      `json:"number_available"`
	Price                        float64                     `json:"price"`
	PurchaseLimit                uint64                      `json:"purchase_limit"`
	DiscountType                 string                      `json:"discount_type"`
	Percentage                   float64                     `json:"percentage"`
	DiscountPrice                float64                     `json:"discount_price"`
	DiscountCode                 string                      `json:"discount_code"`
	AvailableDiscountedTickets   uint64                      `json:"available_discounted_tickets"`
	AdditionalInformationFields  AdditionalInformationFields `gorm:"type:VARCHAR(255)" json:"additional_information_fields,omitempty"`
	IsTransferPaymentFeesToGuest bool
	EventId                      uint64
}

type Event struct {
	ID                 uint64 `id:"EventId" gorm:"primaryKey" json:"id"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	OrganizerID        uint64
	Status             string `json:"status"`
	EventStaffID       uint64 `json:"event_staff_id"`
	TicketID           uint64 `json:"ticket_id"`
	Tickets            []*Ticket
	EventStaff         []*EventStaff
}

type EventStaff struct {
	ID      uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	*User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	EventID uint64
}
