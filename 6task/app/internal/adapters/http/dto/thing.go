package dto

type CreateThingRequest struct {
	Name string `json:"name"`
}

type ThingResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
