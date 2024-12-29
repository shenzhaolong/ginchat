package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	for _, v := range data {
		fmt.Printf("v: %v\n", v)
	}
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "重复密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册",
		})
		return
	}

	if password != repassword || password == "" {
		c.JSON(-1, gin.H{
			"message": "密码有问题",
		})
		return
	}
	user := models.UserBasic{
		Name:          name,
		PassWord:      utils.MakePassword(password, salt),
		Salt:          salt,
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		LogOutTime:    time.Now(),
	}
	db := models.CreateUser(user)
	db.Commit()
	c.JSON(200, gin.H{
		"message": "新增用户成功!",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Printf("update user: %v\n", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改用户不匹配" + err.Error(),
		})
		return
	}

	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})
}

// GetUserList
// @Summary 登陆
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	if !utils.VaildPassword(user.PassWord, user.Salt, pwd) {
		c.JSON(200, gin.H{
			"message": "密码错误",
		})
		return
	}

	data := models.FindUserByNameAndPwd(name, utils.MakePassword(pwd, user.Salt))
	c.JSON(200, gin.H{
		"message": data,
	})
}
