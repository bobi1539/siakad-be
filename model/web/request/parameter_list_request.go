package request

type ParameterListRequest struct {
	Name        string `validate:"required,max=255,min=1" json:"name"`
	ParameterId int64  `validate:"required" json:"parameterId"`
}
