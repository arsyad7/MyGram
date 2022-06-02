package params

type CreateComment struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}
