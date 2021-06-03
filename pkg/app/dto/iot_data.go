package dto

import "time"

type IoTData struct {
	Device string    `json:"device"`
	Value  int       `json:"value"`
	Time   time.Time `json:"time"`
}
