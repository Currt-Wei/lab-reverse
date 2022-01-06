package model

type Reservation struct {
	Id     int    `gorm:"column:id" json:"id"`
	UserId     int    `gorm:"column:user_id" json:"user_id"`
	LabId     int    `gorm:"column:lab_id" json:"lab_id"`
	SeatId     int    `gorm:"column:seat_id" json:"seat_id"`
	ReserveDate	 Time    `gorm:"column:reserve_date" json:"reserve_date"`
	TimeInterval     string    `gorm:"column:time_interval" json:"time_interval"`
	Weekday     int    `gorm:"column:weekday" json:"weekday"`
	Description   string `gorm:"column:description" json:"description"`
}
