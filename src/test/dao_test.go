package test

import (
	"fmt"
	"log"
	"reflect"
	"say/src/dao"
	"say/src/modle"
	"say/src/tools"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

//添加
func Test(t *testing.T) {
	var user modle.User
	db := dao.DB
	db.AutoMigrate(&user)
	log.Println("test")
	user2 := modle.User{
		Name:     "yangteng2",
		Password: "yt1234562",
		Phone:    "1822435610988",
	}
	db.Create(&user2)
	db.First(&user)
	log.Println(user)
}

//登入
func Test2(t *testing.T) {
	phone := "1822435610988"
	DB := dao.DB
	//用户只有一个手机号，根据手机号判断是否有该用户
	var user modle.User
	DB.Debug().Where("phone", phone).Find(&user)
	log.Println(user)
}

//密码加密
func Test3(t *testing.T) {
	s := tools.HashAndSalt([]byte("123456"))
	log.Println(s)
	fmt.Printf("tools.ComparePasswords(\"123456\", []byte(s)): %v\n", tools.ComparePasswords(s, []byte("123456")))
}

//通过反射获取结构体中的字段
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

func Test4(t *testing.T) {
	result := GetFieldName(modle.User{})
	for res := range result {
		fmt.Println(result[res])
	}
}

func Test5(t *testing.T) {
	DB := dao.DB
	var ids []int
	DB.Model(&modle.Education{}).Select("user_id").Where("user_id", 3).Find(&ids)
	var users []modle.User
	var userinfos []modle.UserInfo
	DB.Debug().Where(ids).Find(&users)
	DB.Debug().Where(ids).Find(&userinfos)
}

func Test6(t *testing.T) {
	age := 18
	ageOld := 20
	DB := dao.DB
	var users []modle.User
	var userinfos []modle.UserInfo
	selects := "`user`.`id`,`user`.`name`,`user`.`email`,`user`.`phone`,`user`.`image_url`,user_info.nick_name,user_info.occupation,user_info.exp,user_info.brief,user_info.age,user_info.industry_id,user_info.occupation_id"
	DB.Debug().Select(selects).Joins("left join user_info on user_info.user_id = user.id").Where("age >= ? and age <= ?", age, ageOld).Limit(10).Find(&users).Find(&userinfos)
	for i := 0; i < len(users); i++ {
		log.Println("user::", users[i])
		log.Println("userinfos::", userinfos[i].Age)
	}
}

func getAttributes(arr []modle.UserInfo) []int {
	arrs := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		userInfo := arr[i]
		arrs[i] = userInfo.ID
	}
	return arrs
}

func Test7(t *testing.T) {
	DB := dao.DB
	var users []modle.User
	var userinfos []modle.UserInfo
	DB.Debug().Find(&users)
	DB.Debug().Find(&userinfos)
	ids := getAttributes(userinfos)
	for i := 0; i < len(ids); i++ {
		log.Println("ids::", ids[i])
	}
}

//逻辑删除
func Test8(t *testing.T) {
	DB := dao.DB
	var user modle.User
	user.ID = 3
	DB.Debug().Delete(&user)

	var users []modle.User
	DB.Find(&users)
	for i := 0; i < len(users); i++ {
		log.Println(users[i])
	}
}

//生成token
func Test9(t *testing.T) {
	token, _ := tools.GenToken(1223213)
	log.Println(token)
}

//生成数据库表
func Test10(t *testing.T) {
	DB := dao.DB
	DB.AutoMigrate(&modle.User{}, &modle.Account{})
	DB.AutoMigrate(&modle.User{}, &modle.Additional{})
	DB.AutoMigrate(&modle.User{}, &modle.Comment{})
	DB.AutoMigrate(&modle.User{}, &modle.CommentFace{})
	DB.AutoMigrate(&modle.User{}, &modle.CommentImage{})
	DB.AutoMigrate(&modle.User{}, &modle.Education{})
	DB.AutoMigrate(&modle.User{}, &modle.Fans{})
	DB.AutoMigrate(&modle.User{}, &modle.Friend{})
	DB.AutoMigrate(&modle.User{}, &modle.Hobby{})
	DB.AutoMigrate(&modle.User{}, &modle.Industry{})
	DB.AutoMigrate(&modle.User{}, &modle.Job{})
	DB.AutoMigrate(&modle.User{}, &modle.Jur{})
	DB.AutoMigrate(&modle.User{}, &modle.Recharge{})
	DB.AutoMigrate(&modle.User{}, &modle.Role{})
	DB.AutoMigrate(&modle.User{}, &modle.RoleForJur{})
	DB.AutoMigrate(&modle.User{}, &modle.User{})
	DB.AutoMigrate(&modle.User{}, &modle.UserForRole{})
	DB.AutoMigrate(&modle.User{}, &modle.UserInfo{})
	DB.AutoMigrate(&modle.User{}, &modle.Fans{})
	DB.AutoMigrate(&modle.User{}, &modle.Friend{})
	DB.AutoMigrate(&modle.User{}, &modle.Post{})
	DB.AutoMigrate(&modle.User{}, &modle.PostType{})
}

//Retis测试
func Test11(t *testing.T) {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	// 通过go向redis写入数据 string [key - value]
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 关闭连接
	defer conn.Close()

	// 读取数据 获取名字
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//nameString := r.(string)
	fmt.Println("Manipulate success, the name is", r)
}

//测试redis方法
func Test12(t *testing.T) {
	tools.SetRedisVlue("key1", "1")
	tools.SetRedisVlue("key2", "2")
	value1 := tools.GetRedisVlue("key1")
	log.Print(value1)
}

//测试是否能成功获取到所有的评论数
func Test13(t *testing.T) {
	i := dao.CommentNums(1)
	log.Println("所以总评论数为", i)
}

// 添加几条评论数据
func Test14(t *testing.T) { 
	for i := 0; i < 10; i++ {
		dao.DB.Debug().Create(&modle.Comment{
			UserId:  2,
			BeID:    21,
			Content: fmt.Sprint("测试评论i->",i),
		})
	}
}

// 获取在今天发送的评论数据
func Test15(t *testing.T){
	var comments []modle.Comment
	var now = time.Now()
	dao.DB.Debug().Model(&modle.Comment{}).Where("created_at > ? and created_at < ?", now.AddDate(0,0,-1),now).Find(&comments)
	for i := 0; i < len(comments); i++ {
		log.Println("发表在指定时间之前的评论内容为->",comments[i])
	}
}

// go语言时间操作
func TestTime16(t *testing.T){
	// 获取当前时间
	now := time.Now()
	log.Println("当前时间为->",now)
	// 获取当前时间的年月日
	year, month, day := now.Date()
	log.Println("当前时间的年月日为->",year,month,day)
	// 获取当前时间的时分秒
	hour, min, sec := now.Clock()
	log.Println("当前时间的时分秒为->",hour,min,sec)
	// 获取当前时间的时间戳
	log.Println("当前时间的时间戳为->",now.Unix())
	// 获取当前时间的字符串
	log.Println("当前时间的字符串为->",now.String())
	// 获取当前时间的毫秒数
	log.Println("当前时间的毫秒数为->",now.Nanosecond())
	// 获取昨天
	yesterday := now.AddDate(0,0,-1)
	log.Println("昨天的时间为->",yesterday)
}