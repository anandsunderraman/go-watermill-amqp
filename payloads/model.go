package payloads

type LightMeasured struct {
	Id int `json:"id"`
	Lumens int `json:"lumens"`
	SentAt string `json:"sentAt"`
}
