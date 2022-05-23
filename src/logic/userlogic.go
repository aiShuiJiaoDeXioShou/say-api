package logic

import (
	"errors"
	"say/src/dao"
	"say/src/modle"
	"say/src/tools"
)

//用户登入操作
func UserLogin(user modle.User)(error){
	DB := dao.DB
	//用户只有一个手机号，根据手机号判断是否有该用户
	var _user modle.User
	DB.Debug().Where("phone",user.Phone).Find(&_user)
	if _user.ID == 0 {
		//用户不存在
		return errors.New("用户不存在")
	}
	if !tools.ComparePasswords(_user.Password, []byte(user.Password)) {
		//密码错误
		return errors.New("密码错误")
	}
	return nil
}

//用户注册操作，注册的时候把我们的密码加密一下
func UserRegister(user modle.User)(error){
	DB := dao.DB
	//用户只有一个手机号，根据手机号判断是否有该用户
	var _user modle.User
	DB.Debug().Where("phone",user.Phone).Find(&_user)
	if _user.ID != 0 {
		//用户已存在
		return errors.New("用户已存在")
	}
	//对密码进行加密
	user.Password = tools.HashAndSalt([]byte(user.Password))
	DB.Debug().Create(&user)
	return nil
}

//修改用户密码用户名
func UserUpdate(user modle.User)(error){
	DB := dao.DB
	//用户只有一个手机号，根据手机号判断是否有该用户
	var _user modle.User
	DB.Debug().Where("phone",user.Phone).Find(&_user)
	if _user.ID == 0 {
		//用户不存在
		return errors.New("用户不存在")
	}
	DB.Debug().Model(&_user).Update("name",user.Name)
	DB.Debug().Model(&_user).Update("password",user.Password)
	return nil
}