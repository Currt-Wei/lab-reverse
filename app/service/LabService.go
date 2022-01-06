package service

import "lab-reverse/app/model"

func UpdateLabInfo(lab *model.Lab) error{
	return model.UpdateLabInfo(lab)
}

func AddLab(lab *model.Lab) error{
	return model.AddLab(lab)
}

func DeleteLab(lab *model.Lab) error{
	return model.DeleteLab(lab)
}