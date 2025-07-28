package dto

// User описывает поля, которые не могут быть изменены в ходе контекста запроса
type User struct {
	Id string `json:"id"`
}
