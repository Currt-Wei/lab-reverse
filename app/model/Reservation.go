package model

import (
	"strings"
	"time"
)

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

func RepeatReserve(lab_name, reserve_date, time_interval, account string) (bool,error){
	var reservations []Reservation
	err:=DB.Where("lab_name = ? and reserve_date = ? and time_interval = ?", lab_name, reserve_date, time_interval).Find(&reservations).Error
	if len(reservations)==0{
		return false,err
	}
	for _,reservation := range reservations{
		if reservation.Account==account{
			return true,err
		}
	}
 	return false, err

}

func SearchReserve(account string)(bool,error){
	var reservations []Reservation
	err:=DB.Where("account = ?",account).Find(&reservations).Error
	if err!=nil{
		return false,err
	}
	for _,info:=range reservations{
		interval :=strings.Split(info.TimeInterval,"-")
		t1,_ := time.ParseInLocation("2006-01-02 15:04:05", info.ReserveDate+" "+interval[0],time.Local)
		t2,_ := time.ParseInLocation("2006-01-02 15:04:05", info.ReserveDate+" "+interval[1],time.Local)
		if t1.Before(time.Now())&&t2.After(time.Now()){
			return true,nil
		}
	}
	return false,nil
}