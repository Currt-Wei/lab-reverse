package model

type Reservation struct {
	Id     int    `gorm:"column:id" json:"id"`
	UserId     int    `gorm:"column:user_id" json:"user_id"`
	LabId     int    `gorm:"column:lab_id" json:"lab_id"`
	SeatId     int    `gorm:"column:seat_id" json:"seat_id"`
	ReserveDate	 string    `gorm:"column:reserve_date" json:"reserve_date"`
	TimeInterval     string    `gorm:"column:time_interval" json:"time_interval"`
	Weekday     int    `gorm:"column:weekday" json:"weekday"`
	Description   string `gorm:"column:description" json:"description"`
}

func (R Reservation) TableName() string {
	return "reservation"
}

func ReserveSeat(reservation *Reservation) error{
	return DB.Create(&reservation).Error
}

func SearchSeat(date, time string, labId int) ([]Reservation,error) {
	var reservations []Reservation
	err := DB.Where("lab_id = ? and reserve_date = ? and time_interval = ?", labId, date, time).Find(&reservations).Error
	return reservations,err
}