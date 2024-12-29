package models

import "time"

type Job struct {
	UserId          string    `json:"user_id,omitempty"`
	CompanyName     string    `json:"company_name"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ApplicationDate time.Time `json:"application_date"`
	Status          string    `json:"status"`
}
