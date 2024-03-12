package models

type Book struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	Tittle      string `gorm:"type:varchar(300)" json:"tittle"`
	Description string `gorm:"type:text(300)" json:"description"`
	Aothir      string `gorm:"type:varchar(300)" json:"author"`
	PublishDate string `gorm:"type:date(300)" json:"publish_date"`
}
