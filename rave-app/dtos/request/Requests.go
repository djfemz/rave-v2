package dtos

import "github.com/djfemz/rave/rave-app/models"

type AuthRequest struct {
	CreateUserRequest
}

type CreateEventRequest struct {
	Name               string `json:"name"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	OrganizerId        uint64 `json:"organizer_id"`
}

type UpdateEventRequest struct {
	Name               string `json:"name"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	OrganizerId        uint64 `json:"organizer_id"`
}

type CreateTicketRequest struct {
	Type                         string                             `json:"ticket_type"`
	Name                         string                             `json:"name"`
	Stock                        uint64                             `json:"stock"`
	NumberAvailable              uint64                             `json:"number_in_stock"`
	Price                        float64                            `json:"price"`
	PurchaseLimit                uint64                             `json:"purchase_limit"`
	DiscountType                 string                             `json:"discount_type"`
	Percentage                   float64                            `json:"percentage"`
	DiscountPrice                float64                            `json:"discount_price"`
	DiscountCode                 string                             `json:"discount_code"`
	AvailableDiscountedTickets   uint64                             `json:"available_discounted_tickets"`
	IsTransferPaymentFeesToGuest bool                               `json:"is_transfer_payment_fees_to_guest"`
	AdditionalInformationFields  models.AdditionalInformationFields `json:"additional_information_fields"`
	EventId                      uint64                             `json:"event_id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
}

type AddEventStaffRequest struct {
	StaffEmails []string `json:"staff_emails"`
	EventId     uint64   `json:"event_id"`
}

type CreateEventStaffRequest struct {
	StaffEmails []string `json:"staff_emails"`
	EventId     uint64   `json:"event_id"`
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailNotificationRequest struct {
	Sender     Sender      `json:"sender"`
	Recipients []Recipient `json:"to"`
	Subject    string      `json:"subject"`
	Content    string      `json:"htmlContent"`
}

func NewEmailNotificationRequest(recipient, content string) *EmailNotificationRequest {
	return &EmailNotificationRequest{
		Sender: Sender{
			Email: "noreply@email.com",
			Name:  "rave",
		},
		Recipients: []Recipient{
			{Email: recipient, Name: "Friend"},
		},
		Subject: "rave",
		Content: content,
	}
}
