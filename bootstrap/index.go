package bootstrap

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func IndexTest() {

	// 创建一个 gorm.DB 类型的变量
	var db *gorm.DB
	// 调用 Open 方法，传入驱动名和连接字符串
	db, err := gorm.Open(sqlite.Open("root:123456@./test.db"), &gorm.Config{})
	// 检查是否有错误
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	// 打印成功信息
	fmt.Println("连接数据库成功")

	// 关闭GORM DB实例
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get *sql.DB")
	}
	defer sqlDB.Close() // 关闭数据库连接
}
