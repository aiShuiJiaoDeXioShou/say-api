package dao

import (
	"log"
	"say/src/modle"

	"github.com/gin-gonic/gin"
)

func init(){
	log.Println("dao -> 用户mapper,被调用了")
}

//查询用户信息
func UserInfo(userid int) (modle.UserInfo,modle.User) {
	//用户扩展信息
	var userinfo modle.UserInfo
	//用户基本信息
	var user modle.User
	DB.Debug().Where("user_id = ?",userid).Find(userinfo)
	DB.Debug().Where("id = ?",userid).Find(user)
	return userinfo,user
}

//查询用户职业
func UserOccupation(userid int) modle.Job {
	var userinfo modle.UserInfo
	DB.Debug().Where("user_id = ?",userid).Find(&userinfo)
	var job modle.Job
	DB.Debug().Find(&job,userinfo.Occupation)
	return job
}

//查询用户行业
func UserIndustry(userid int) modle.Industry {
	var userinfo modle.UserInfo
	DB.Debug().Where("user_id = ?",userid).Find(&userinfo)
	var industry modle.Industry
	DB.Debug().Find(&industry,userinfo.IndustryID)
	return industry
}

//查询用户兴趣爱好
func UserHobby(userid int) []modle.Hobby {
	var hobbies []modle.Hobby
	DB.Debug().Where("user_id = ?",userid).Find(&hobbies)
	return hobbies
}

//查询指定用户的学历
func EducationUser(user modle.User) []modle.Education{
	var educations []modle.Education
	DB.Where("user_id",user.ID).Find(&educations)
	return educations
}

//分类,根据行业查询用户
func CategoryListIndustry(industry modle.Industry,pages ...int) gin.H {
	//如果没有指定查询条数，默认为10条查询结果
	var _users []modle.User
	//拿到用户信息
	var _userInfo []modle.UserInfo
	DB.Debug().Where("industry_id = ?", industry.ID).Limit(10).Offset(getPage(pages...)).Find(&_userInfo)
	//从用户详细信息中取用户信息
	DB.Debug().Where("id in (?)", getAttributes(_userInfo)).Limit(10).Offset(getPage(pages...)).Find(&_users)
	data := gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"user_data":      _users,
			"user_info_data": _userInfo,
		},
		"success": true,
	}
	return data
}

//分类,根据学历查询用户
func CategoryListEducation(education modle.Education,pages ...int) gin.H {
	var ids []int
	DB.Model(&modle.Education{}).Select("user_id").Limit(10).Offset(getPage(pages...)).Where("user_id",3).Find(&ids)
	var users []modle.User
	var userinfos []modle.UserInfo
	DB.Debug().Where(ids).Find(&users) 
	DB.Debug().Where(ids).Find(&userinfos)
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"user_data":      users,
			"user_info_data": userinfos,
		},
		"success": true,
	}
}

//分类,根据职业查询用户
func CategoryListOccupation(occupation modle.Industry)([]modle.User,[]modle.UserInfo) {
	var userinfos []modle.UserInfo
	DB.Debug().Where("occupation_id = ?", occupation.ID).Find(&userinfos)
	var users []modle.User
	DB.Debug().Where("id in (?)", getAttributes(userinfos)).Find(&users)
	return users,userinfos
}

//分类,根据兴趣爱好查询用户
func CategoryListInterest(hobby modle.Hobby,pages ...int) {
	var ids []int
	var page int
	if len(pages) == 0 {
		page = 0
	}else {
		page = pages[0]
	}
	DB.Model(&modle.Hobby{}).Select("user_id").Limit(10).Offset(page).Where("user_id",3).Find(&ids)
}

//分类根据年龄查询用户
func CategoryListAge(age int,ageOld int,pages ...int)([]modle.User,[]modle.UserInfo) {
	var users []modle.User
	var userinfos []modle.UserInfo
	selects := "`user`.`id`,`user`.`name`,`user`.`email`,`user`.`phone`,`user`.`image_url`,user_info.nick_name,user_info.occupation,user_info.exp,user_info.brief,user_info.age,user_info.industry_id,user_info.occupation_id"
	DB.Debug().Select(selects).Joins("left join user_info on user_info.user_id = user.id").Where("age >= ? and age <= ?", age,ageOld).Limit(10).Offset(getPage(pages...)).Find(&users).Find(&userinfos)
	return users,userinfos
}

//删除全部用户信息
func DeleteAllUser(userid int) {
	DB.Debug().Delete(modle.User{},userid)
	DB.Debug().Where("user_id = ?",userid).Delete(&modle.UserInfo{})
	DB.Debug().Where("user_id = ?",userid).Delete(&modle.Education{})
	DB.Debug().Where("user_id = ?",userid).Delete(&modle.Hobby{})
}

//查询一个用户拥有的角色
func UserHaveRole(userid int) []modle.Role {
	var roleids []int
	var roles []modle.Role
	DB.Debug().Model(&modle.UserForRole{}).Select("role_id").Where("user_id = ?",userid).Find(&roleids)
	DB.Debug().Where("id in (?)",roleids).Find(&roles)
	return roles
}

//查询一个角色拥有的权限
func RoleForJur(roleid int) []modle.Jur {
	var jurids []int
	var jurs []modle.Jur
	DB.Debug().Model(&modle.RoleForJur{}).Select("jur_id").Where("user_id = ?",roleid).Find(&jurids)
	DB.Debug().Where("id in (?)",jurids).Find(&jurs)
	return jurs
}

//查询一个用户是否拥有该权限
func IsUserHaveJur(userid int,jurid int) bool {
	//通过userid拿到全部角色
	ufrs := UserHaveRole(userid)

	//通过全部角色拿到全部权限,去除相同权限
	for _,ufr := range ufrs {
		
		jurs := RoleForJur(ufr.ID)

		//判断jur是否在jurs中
		for _,jur := range jurs {
			if jur.ID == jurid {
				return true
			}
		}

	}
	return false
}

//查询指定权限
func GetJur(jurid int) modle.Jur {
	var jur modle.Jur
	DB.Debug().Where("id = ?",jurid).Find(&jur)
	return jur
}

//查询一个用户所有的权限
func UserJur(userid int) []modle.Jur {
	var roleid []int
	var jurids []int
	var jurs []modle.Jur
	DB.Debug().Model(&modle.UserForRole{}).Select("role_id").Where("user_id = ?",userid).Find(&roleid)
	DB.Debug().Model(&modle.RoleForJur{}).Distinct().Select("jur_id").Where("role_id in (?)",roleid).Find(&jurids)
	DB.Debug().Where("id in (?)",jurids).Find(&jurs)
	return jurs
}