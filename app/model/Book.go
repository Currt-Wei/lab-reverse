package model

type Book struct {
	Id        int       `gorm:"column:id" json:"id"`
	UserId    string    `gorm:"column:user_id" json:"user_id"`
	LabId     string    `gorm:"column:lab_id" json:"lab_id"`
	SeatId    string    `gorm:"column:seat_id" json:"seat_id"`
	StartTime LocalTime `gorm:"column:start_time" json:"start_time"`
	EndTime   LocalTime `gorm:"column:end_time" json:"end_time"`

}

func (B Book) TableName() string {
	return "seats"
}

func ReverseSeat(book Book) error {
	return DB.Create(&book).Error
}
