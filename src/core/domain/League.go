package domain

type League struct {
	Id          string
	DisplayName string
	OwnerId     string
	SecureId    string
	IsActive    bool
	IsRegOpen   bool
}
