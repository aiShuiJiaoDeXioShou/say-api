package dao

import (
	"log"
	"say/src/modle"
)

// 这个用于操作用户充值与用户账户的dao层
func init(){
	log.Println("dao -> 充值,被调用了")
}

// 创建用户账户
func CreateAccount(account *modle.Account) bool{
	res := DB.Debug().Create(&account)

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	return true
}

// 根据充值方案进行用户充值
func UserRechargeByProgramme(rechargeProgrammeID int,userID int64) bool{
	// 查询用户的充值方案
	rechargeProgrammeInfo := &modle.RechargeProgramme{}
	res := DB.Debug().Where("id = ?",rechargeProgrammeID).Find(rechargeProgrammeInfo)
	// 根据用户充值方案进行
	UserRecharge(&modle.Recharge{
		UserID:    userID,
		Money:     rechargeProgrammeInfo.Money,
		TypeID:    rechargeProgrammeInfo.TypeID,
	})

	// 根据赠送ID进行赠送处理
	if rechargeProgrammeInfo.AdditionalID != 0 {
		resAdd := AdditionalHandle(rechargeProgrammeInfo.AdditionalID,userID)
		if resAdd == nil {
			log.Println("赠送失败")
			return false
		}
	}

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	return true
}

// 用户充值
func UserRecharge(recharge *modle.Recharge) bool{
	res := DB.Debug().Create(recharge)
	// 充值之后更新用户账户信息
	updateRes := DB.Debug().Model(&modle.Account{}).Where("user_id = ?",recharge.UserID).Update("balance",SelectAccountInfo(recharge.UserID).Money+recharge.Money)

	if res.Error != nil {
		log.Println("用户充值失败,err->%v",res.Error)
		return false
	}
	if updateRes.Error != nil {
		log.Println("用户更新账户余额异常,err->%v",updateRes.Error)
		return false
	}
	return true
}


// 查询用户账户信息
func SelectAccountInfo(userID int64) *modle.Recharge{
	recharge := &modle.Recharge{}
	res := DB.Debug().Where("user_id = ?",userID).Find(recharge)
	if res.Error != nil {
		log.Println("查询用户失败->%v",res.Error)
		return nil
	}
	return recharge
}

// 赠送处理Mapper
func AdditionalHandle(additionalID int64,userID int64) *modle.Additional{
	additional := &modle.Additional{}
	res := DB.Debug().Where("id = ?",additionalID).Find(additional)
	if additional.Money != 0{
		UserRecharge(&modle.Recharge{
			UserID:    userID,
			Money:     additional.Money,
			TypeID:    2,
		})
	}

	if res.Error != nil {
		log.Println(res.Error)
		return nil
	}
	return additional
}