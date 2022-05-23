package test

import (
	"log"
	"say/src/dao"

	// "say/src/modle"
	"testing"
)

func Test1(t *testing.T) {
	/* dao.DB.Debug().Create(&modle.Jur{
		JurName: "管理员",
		Target:  "/admin",
		Effect:  "进去管理员页面",
	}) */
	jur := dao.GetJur(1)
	log.Println(jur)
}
