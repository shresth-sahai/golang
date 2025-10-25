package models

import "time"

type URLMapping struct {
	ID string  `json:"id"`
	ShortCode string `json:"short_code"`
	OrignalURL string `json:"orignal_url"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	Clicks int `json:"clicks"`
	CustomCode bool `json:"custom_code"`
}

func (u *URLMapping) isExpired () bool {
	if u.ExpiresAt == nil {
		return false
	}
	return  time.Now().After(*u.ExpiresAt)
}

func (u *URLMapping) IncrementClicks(){
	u.Clicks++;
}