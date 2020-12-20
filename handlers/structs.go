package handlers

type ImageResponse struct{
	ID string `json:"id"`
	Status int `json:"status"`
	Message string `json:"message"`
}

type RetrieveImages struct {
	ImageName string `json:"image_name"`
	ImagePath string `json:"image_path"`
}
