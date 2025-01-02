package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:Aa772639270!@tcp(172.25.82.123:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})

	// // Create
	// user := &models.UserBasic{
	// 	Name:          "沈朝龙",
	// 	LoginTime:     time.Now(),
	// 	HeartbeatTime: time.Now(),
	// 	LogOutTime:    time.Now(),
	// }
	// db.Create(user)

	// fmt.Printf("db.First(user, 1): %v\n", db.First(user, 1))

	// // Read
	// db.Model(user).Update("PassWord", "1234")

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
