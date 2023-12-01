package models

type Tool struct {
	ID          uint   `json:"id" gorm:""primaryKey"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags" gorm:"foreignKey:Tag;"`
}

type Tag struct {
	Tag string `gorm:"type:varchar[]"`
}
