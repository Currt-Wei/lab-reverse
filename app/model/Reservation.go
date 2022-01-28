package model

type Reservation struct {
	Id     int    `gorm:"column:id" json:"id"`
	UserName	string    `gorm:"column:user_name" json:"user_name"`
	Account     string    `gorm:"column:account" json:"account"`
	LabName     string    `gorm:"column:lab_name" json:"lab_name"`
	SeatName    string    `gorm:"column:seat_name" json:"seat_name"`
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

func SearchSeat(date, time,labName string) ([]Reservation,error) {
	var reservations []Reservation
	err := DB.Where("lab_name = ? and reserve_date = ? and time_interval = ?", labName, date, time).Find(&reservations).Error
	return reservations,err
}

func DeleteReserve(reservation *Reservation) error{
	return DB.Where("account=? and lab_name=? and seat_name = ? and reserve_date=? and time_interval=?",reservation.Account,reservation.LabName,reservation.SeatName,reservation.ReserveDate,reservation.TimeInterval).Delete(&reservation).Error
}

func RepeatReserve(lab_name, reserve_date, time_interval string) (bool,error){
	var reservation *Reservation
	err:=DB.Where("lab_name = ? and reserve_date = ? and time_interval = ?", lab_name, reserve_date, time_interval).Find(&reservation).Error
	if reservation.Account!=""{
		return true,err
	}
	return false, err

}