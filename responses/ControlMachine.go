package responses

import "go-clean/models"

type MachineDefault struct {
	Status string `json:"status" `
	Msg    string `json:"msg" `
}

type MachineData struct {
	Status string         `json:"status" `
	Data   models.Machine `json:"data" `
}
