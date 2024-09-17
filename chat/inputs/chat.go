package inputs

type CreateChatInput struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	ServerID string `json:"serverId" validate:"required"`
}
