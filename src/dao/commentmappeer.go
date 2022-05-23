package dao

import (
	"log"
	"say/src/modle"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("dao -> 用户聊天mapper,被调用了")
}

// 获取指定ID的子评论数
func GetComments(id int) (gin.H, error) {
	//这个是主评论
	comment := &modle.Comment{}
	res := DB.Debug().Model(&modle.Comment{}).Where("id",id).Find(comment)
	if res.Error != nil {
		return nil, res.Error
	}

	//这个是子评论
	commentchildrens := &[]modle.Comment{}
	reschild := DB.Debug().Model(&modle.Comment{}).Where("be_id",comment.ID).Find(commentchildrens)
	if reschild.Error != nil {
		return nil, reschild.Error
	}

	return gin.H{
		"comment" : comment,
		"commentchildrens" : commentchildrens,
	},nil
}
// 查看最近的评论(一天以内)
func GetLatelyComments(){
	var comments []modle.Comment
	var now = time.Now()
	DB.Debug().Model(&modle.Comment{}).Where("created_at > ? and created_at < ?", now.AddDate(0,0,-1),now).Find(&comments)
}

// 查看热度最高的评论
func GetHotMaxComments(){
	
}

// 获取按时间搜索前30的评论
func GetTimeSearchComments(){
	// 根据created_at时间排名评论数据
	var comments []modle.Comment
	DB.Debug().Model(&modle.Comment{}).Order("created_at desc").Limit(30).Find(&comments)
}

var columnNums int
func CommentNums(id int64)int {
	commentNums(id)
	return columnNums
}

// 获取指定ID下所有的评论条目
func commentNums(id int64)[]int64{
	var ids []int64
	DB.Debug().Model(&modle.Comment{}).Select("id").Where("be_id",id).Find(&ids)
	columnNums += len(ids)
	if len(ids)>0{
		for i := 0; i < len(ids); i++ {
			commentNums(ids[i])
		}
	}
	return ids
}