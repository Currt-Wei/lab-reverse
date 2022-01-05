package service

import "lab-reverse/app/model"

func AddSeat(seat model.Seat) error{
	err:=model.AddSeat(seat)
	return err
}
