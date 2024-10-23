package todo

type User struct {
	// json:"-": это поле не должно быть включено в JSON-ответ при серилизации
	// db:"id": указывает, что это поле связано с колонкой id в таблице базы данных.
	// binding:"required": поле обязательное для заполнения.
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
