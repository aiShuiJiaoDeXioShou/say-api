package modle

import (
	"time"

	"gorm.io/gorm"
)

//账户的信息表
type Account struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserId    int    `gorm:"column:user_id"`
	AccountName   string `gorm:"type:varchar(100)"` //账户名
	Balance float64 `gorm:"type:decimal(10,2)"`//账户余额
	BankCard string `gorm:"type:varchar(100)"` //银行卡号
	UserID string `gorm:"type:varchar(100)"` //用户身份证号
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//充值方案表
type RechargeProgramme struct {
	ID        int64   `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	Name      string  `json:"name"` //充值方案名称
	Money     float64 `json:"money"` //充值金额
	TypeID    int64   `json:"type"` //充值类型id
	AdditionalID int64   `json:"additional_id"` //这个是赠送的东东id
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//充值信息表
type Recharge struct {
	ID        int64   `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserID    int64   `json:"user_id"` //用户ID
	Money     float64 `json:"money"` //充值金额
	TypeID    int64   `json:"type"` //充值类型id
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//这个是充值之后赠送的东西
type Additional struct {
	ID        int64   `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	RechargeProgrammeID int64 `json:"recharge_programme_id"` //充值方案id
	JurID int64//赠送的权限ID
	Money     float64 `json:"money"`//赠送的金额
	Integral  float64 `json:"integral"`//赠送的积分
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//充值类型表
type RechargeType struct {
	ID        int64  `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	Name      string `json:"name"`
	Describe  string `json:"describe"`//对充值信息的描述
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}