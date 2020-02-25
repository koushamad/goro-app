package resource

type HelloRequest struct {
	Name     string `json:"name" xml:"name" binding:"required"`
}
