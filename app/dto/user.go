package dto

import "lab-reverse/app/model"

type UserDto struct {
	Id uint 		`json:"id"`
	Account string	`json:"account"`
	Name string		`json:"name"`
	Email string	`json:"email"`
	Telephone string	`json:"telephone"`
	College string	`json:"college"`
	Degree string	`json:"degree"`
	Grade string	`json:"grade"`
	Identity string	`json:"identity"`
	Enable int		`json:"enable"`

	RoleName string	`json:"role_name"`
}

type UserDtoAndRole struct {
	UserDto
	Role	model.Role		`json:"role"`
	Roles	[]model.Role	`json:"roles"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id: user.Id,
		Account: user.Account,
		Name: user.Name,
		Email: user.Email,
		Telephone: user.Telephone,
		College: user.College,
		Degree: user.Degree,
		Grade: user.Grade,
		Identity: user.Identity,
		Enable: user.Enable,
	}
}

func ToUserDtos(users []model.User) []UserDto {
	var dtos []UserDto
	for _, user := range users {
		dtos = append(dtos, ToUserDto(user))
	}
	return dtos
}

func ToUserDtoAndRole(user model.User) UserDtoAndRole {
	return UserDtoAndRole {
		UserDto: UserDto {
			Id:        user.Id,
			Account:   user.Account,
			Name:      user.Name,
			Email:     user.Email,
			Telephone: user.Telephone,
			College:   user.College,
			Degree:    user.Degree,
			Grade:     user.Grade,
			Identity:  user.Identity,
			Enable:    user.Enable,
		},
		Role: user.Role,
		Roles: user.Roles,
	}
}

func ToUserDtoAndRoles(users []model.User) []UserDtoAndRole {
	var dtos []UserDtoAndRole
	for _, user := range users {
		dtos = append(dtos, ToUserDtoAndRole(user))
	}
	return dtos
}