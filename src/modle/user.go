package modle

import (
	"time"

	"gorm.io/gorm"
)

//用户表
type User struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	Name      string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"type:varchar(100)"`
	ImageUrl  string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//用户基本信息表
type UserInfo struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserId    int    `gorm:"column:user_id"`
	NickName  string `gorm:"type:varchar(100)"` //昵称
	Occupation string `gorm:"type:varchar(100)"` //工作单位
	Exp int //经验
	Brief   string `gorm:"type:varchar(100)"` //简介
	Age    int    `gorm:"type:int"` //年龄
	IndustryID int `gorm:"type:int"` //行业
	OccupationID int `gorm:"type:int"` //职业id
	AddressID int `gorm:"type:int"` //地址id
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//学历表
type Education struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	Grade     int   `gorm:"column:grade"` //学历等级，幼儿园0，小学1，初中2，高中3，大学4，研究生5，博士6，博士后7
	UserId    int    `gorm:"column:user_id"`
	School    string `gorm:"type:varchar(100)"` //学校
	Major     string `gorm:"type:varchar(100)"` //专业
	StartTime string `gorm:"type:varchar(100)"` //开始时间
	EndTime   string `gorm:"type:varchar(100)"` //结束时间
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//用户兴趣爱好表
type Hobby struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	Hobby string `gorm:"type:varchar(100)"`
	UserID int `gorm:"column:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//行业
type Industry struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	IndustryName string `gorm:"type:varchar(100)"` //行业的名称
	Details string `gorm:"type:varchar(100)"` //行业的详情
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//职业
type Job struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	IndustryID int `gorm:"column:industry_id"` //行业ID
	JobName string `gorm:"type:varchar(100)"` //职业名称
	Details string `gorm:"type:varchar(100)"` //职业的详情
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//一个用户拥有很多角色
type UserForRole struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserID int `gorm:"column:user_id"` //用户ID
	RoleID int `gorm:"column:role_id"` //角色ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//角色表
type Role struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	RoleName  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//一个角色拥有很多权限
type RoleForJur struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	RoleID int `gorm:"column:role_id"` //角色ID
	JurID int `gorm:"column:jur_id"` //权限ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//权限表
type Jur struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	JurName   string `gorm:"type:varchar(100)"` //权限名称
	Target	  string `gorm:"type:varchar(100)"` //权限目标
	Effect    string `gorm:"type:varchar(100)"` //此权限的效果
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}