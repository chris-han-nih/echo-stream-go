package models

type App struct {
	Id        int64  `db:"id" json:"id"`
	UserId    int64  `db:"user_id" json:"user_id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Token     string `db:"token" json:"token"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
