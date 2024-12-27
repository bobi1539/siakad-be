package request

type ParameterRequest struct {
	Name        string `validate:"required,max=255,min=1" json:"name"`
	Description string `validate:"required,max=255,min=1" json:"description"`
}
