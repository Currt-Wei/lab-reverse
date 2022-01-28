package model

type Seat struct {
	Id        int       `gorm:"column:id" json:"id"`
	LabName     string    `gorm:"column:lab_name" json:"lab_name"`
	SeatId    int    `gorm:"column:seat_id" json:"seat_id"`
	SeatName    string    `gorm:"column:seat_name" json:"seat_name"`
}

func (S Seat) TableName() string {
	return "seats"
}

func AddSeat(Seat Seat) error{
	return DB.Create(&Seat).Error
}

func FindSeatBySeatId(seatId int) (Seat,error){
	var seat *Seat
	err:=DB.Where("seat_id = ? ",seatId).Find(&seat).Error
	return *seat,err
}

func FindSeatBySeatNameAndLabName(seatName, labName string) (Seat,error){
	var seat *Seat
	err:=DB.Where("seat_name = ? and lab_name = ?",seatName, labName).Find(&seat).Error
	return *seat,err
}