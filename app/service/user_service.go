package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lab-reverse/app/model"
	"lab-reverse/app/model/request"
	"lab-reverse/global"
	"reflect"
)

type UserService struct {}

// Register 注册
func (u *UserService) Register(user *model.User) (error, *model.User) {
	// 判断学号是否重复
	var temp model.User
	global.DB.Where("account", user.Account).Find(&temp)
	if temp.Id != 0 {
		return errors.New("该用户已注册"), nil
	}

	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败"), nil
	}

	// 替换密码
	user.Password = string(password)
	err = global.DB.Create(user).Error

	// 添加该用户-角色到关联表
	global.DB.Create(&model.UserRole{
		UserId: user.Id,
		RoleId: 666666,
	})

	return err, user
}

// Login 登录
func (u *UserService) Login(loginUser *model.User) (error, *model.User) {
	// 查出当前用户
	var temp model.User
	global.DB.Preload("Role").Preload("Roles").Where("account", loginUser.Account).Find(&temp)
	if temp.Id == 0 {
		// 查不到该用户，返回错误
		return errors.New("该用户未注册"), nil
	}
	// 对比密码是否错误
	err := bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(loginUser.Password))
	if err != nil {
		// 密码错误
		return errors.New("账号或密码错误"), nil
	}
	return nil, &temp
}

// UpdateUserInfo 更新用户信息
func (u *UserService) UpdateUserInfo(user model.User) (error, model.User) {
	err := global.DB.Updates(&user).Error
	global.DB.Find(&user)
	return err, user
}

// UpdatePassword 更新密码
func (u *UserService) UpdatePassword(user *model.User, newPassword string) (error, *model.User) {
	// 先查出用户的密码
	global.DB.First(&user)
	old, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(old)
	// 校对源密码是否一致
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newPassword))
	if err != nil {
		return errors.New("源密码错误"), nil
	}

	// 加密密码，存入数据库
	passwd, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密错误"), nil
	}
	user.Password = string(passwd)
	err = global.DB.Select("password").Updates(user).Error
	return err, user
}

// GetUserList 获取用户列表
func (u *UserService) GetUserList(pageInfo request.PageInfo, userFilter request.UserFilter) (error, []model.User, int64) {
	var users []model.User
	offset := (pageInfo.Page-1) * pageInfo.Limit
	db := global.DB
	// 添加筛选条件
	// 根据反射获取结构体中的每个值
	value := reflect.TypeOf(userFilter)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Name
		v := reflect.ValueOf(userFilter).Field(i).String()
		if v != "" {
			v = "%" + v + "%"
			db = db.Where(field + " LIKE ?" , v)
		}
	}

	// 计算总条数
	var total int64	// 总条数
	err := db.Debug().Table("user").Count(&total).Error
	if err != nil {
		return errors.New("计算总条数出错"), nil, 0
	}

	// 查找所有的user
	err = db.Limit(pageInfo.Limit).Offset(offset).Preload("Role").Preload("Roles").Find(&users).Error
	if err != nil {
		return errors.New("管理员查询所有用户失败"), nil, 0
	}

	return nil, users, total
}

// GetSelfInfo 获取一个人的信息
func (u *UserService) GetSelfInfo(id uint) (model.User, error) {
	var user = model.User{Id: id}
	err := global.DB.Preload("Role").Preload("Roles").First(&user).Error
	return user, err
}

// UpdateUserStatus 更新用户状态
func (u *UserService) UpdateUserStatus(id uint, enable int) {
	global.DB.Model(&model.User{}).Select("enable").
		Where("id", id).Updates(&model.User{Id: id, Enable: enable})
}