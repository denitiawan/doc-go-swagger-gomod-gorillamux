package auth

type TokenDto struct {
	Key   string `validate:"required,min=1,max=200" json:"key"`
	Value string `validate:"required,min=1,max=200" json:"value"`
}

func NewTokenDto(key string, value string) *TokenDto {
	return &TokenDto{
		Key:   key,
		Value: value,
	}
}
