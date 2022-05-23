package modle

import (
	"time"

	"gorm.io/gorm"
)

// 帖子表
type Post struct {
	gorm.Model
	UserID    int64  `json:"user_id"`    //用户ID
	Title     string `json:"title"`      //帖子标题
	Content   []byte `json:"content"`    //帖子内容
	Type      int64  `json:"type"`       //帖子类型
	ClassType int64  `json:"class_type"` //帖子分类
}

// 帖子的类型
type PostType struct {
	gorm.Model
	Name       string `json:"name"`        //帖子类型名称
	Describe   string `json:"describe"`    //帖子类型描述
	ChildrenID int64  `json:"children_id"` //帖子类型的子类型 ,如果是父类型就是0
}

// 评论表
type Comment struct {
	ID        int    `gorm:"primary_key" json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserId    int    `json:"user_id"`                  //评论的用户
	Content   string `json:"content" gorm:"type:text"` //评论内容
	BeID      int    //被评论的ID,如果是顶楼的话,则ID为零
	PostsID   int64  `json:"posts_id" gorm:type:int`
	GroupID   int    `json:"comment_id" gorm:"type:int"` //评论的组ID
	Give      int64  `json:"give" gorm:"give"`           //点赞量
	Forward   int64  `json:"forward" gorm:"forward"`     //转发量
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// 评论图片的位置
type CommentImage struct {
	ID        int    `json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id"`
	ImageUrl  string `json:"image_url" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//评论的表情包
type CommentFace struct {
	ID        int    `json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id"`
	FaceUrl   string `json:"face_url" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//好友
type Friend struct {
	ID        int `json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserId    int `json:"user_id" gorm:"column:user_id"`
	FriendId  int `json:"friend_id" gorm:"column:friend_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//粉丝
type Fans struct {
	ID        int `json:"id" gorm:"column:id" form:"id" query:"id" AutoMigrate:"true"`
	UserId    int `json:"user_id" gorm:"column:user_id"`
	FansId    int `json:"fans_id" gorm:"column:fans_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
