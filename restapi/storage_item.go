package restapi

type StorageItem struct {
	ID string `json:"id"`
	Password string `json:"-"`
	Name string `json:"name"`
	EMail string `json:"email"`
}
