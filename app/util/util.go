package util

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var PublicPath string

//获取当前登录管理员id
func CurrentAdminId(c *gin.Context) string {
	userId, _ := c.Get("user_id")
	return userId.(string)
}

//获取当前用户信息
func CurrentUser(c *gin.Context) *model.User {
	u := &model.User{}
	userId, _ := c.Get("user_id")
	cmf.NewDb().First(u, "id = ?", userId)
	return u
}

type role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 获取当前用户角色
func CurrentRole(c *gin.Context) []role {
	userId, _ := c.Get("user_id")
	userIdInt, _ := strconv.Atoi(userId.(string))
	return GetRoleById(userIdInt)
}

// 根据用户id获取所有角色
func GetRoleById(userId int) []role {
	var result []role
	prefix := cmf.Conf().Database.Prefix
	cmf.NewDb().Table(prefix+"role_user ru").Select("r.id,r.name").
		Joins("INNER JOIN "+prefix+"role r ON ru.role_id = r.id").
		Where("user_id = ?", userId).
		Scan(&result)
	return result
}

// 是否超级管理员
func SuperRole(c *gin.Context, t int) bool {

	type resultStruct struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var result []resultStruct
	userId, _ := c.Get("user_id")

	if userId == "1" {
		return true
	}

	prefix := cmf.Conf().Database.Prefix
	cmf.NewDb().Table(prefix+"role_user ru").Select("r.id,r.name").
		Joins("INNER JOIN "+prefix+"role r ON ru.role_id = r.id").
		Where("ru.user_id = ?", userId).
		Scan(&result)
	for _, v := range result {
		if v.Id == t {
			return true
		}
	}
	return false
}

//获取网站上传配置信息
func UploadSetting(c *gin.Context) *model.UploadSetting {
	session := sessions.Default(c)
	uploadSettingStr := session.Get("uploadSetting")
	option := &model.Option{}
	uploadSetting := &model.UploadSetting{}
	if uploadSettingStr == nil {
		uploadResult := cmf.NewDb().First(option, "option_name = ?", "upload_setting") // 查询
		if uploadResult.RowsAffected > 0 {
			uploadSettingStr = option.OptionValue
			//存入session
			session.Set("uploadSetting", uploadSettingStr)
		}
	}

	//读取的数据为json格式，需要进行解码
	_ = json.Unmarshal([]byte(uploadSettingStr.(string)), uploadSetting)
	return uploadSetting
}

//添加用户操作日志
func SetLog(c *gin.Context, module string, controller string, action string, message string) {
	currentUser := CurrentUser(c)
	cmf.NewDb().Create(&model.Log{
		ModuleName:     module,
		ControllerName: controller,
		ActionName:     action,
		Url:            c.Request.URL.String(),
		RequestIp:      c.ClientIP(),
		UserId:         currentUser.Id,
		UserNickname:   currentUser.UserNickname,
		Message:        message,
		CreateAt:       time.Now().Unix(),
	})
}

// 获取真实路径
func CurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取真实url
func GetFileUrl(path string) string {

	if path == "" {
		return ""
	}

	domain := cmf.Conf().App.Domain
	prevPath := domain + "/uploads/" + path
	return prevPath
}

// 去除空格回车
func TrimAll(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}

// 获取数据库配置信息
func SiteSettings() map[string]interface{} {
	option := &model.Option{}
	cmf.NewDb().First(option, "option_name = ?", "site_info") // 查询

	m := make(map[string]interface{}, 5)
	err := json.Unmarshal([]byte(option.OptionValue), &m) //第二个参数要地址传递
	if err != nil {
		return m
	}
	return nil
}

func ToLowerInArray(search string, arr []string) bool {
	for _, item := range arr {
		if strings.ToLower(search) == strings.ToLower(item) {
			return true
		}
	}
	return false
}

func AuthAccess(c *gin.Context) []model.AuthAccessRule {

	session := sessions.Default(c)
	res := session.Get("authAccessRule")
	if session.Get("authAccessRule") == nil {

		user := CurrentUser(c)
		// 获取当前用户全部的权限列表
		role := GetRoleById(user.Id)
		var roleIds []string
		for _, v := range role {
			roleIds = append(roleIds, strconv.Itoa(v.Id))
		}

		roleIdsStr := strings.Join(roleIds, ",")

		query := []string{"1=1"}
		var queryArgs []interface{}

		if roleIdsStr != "" {
			query = append(query, "role_id in (?)")
			queryArgs = append(queryArgs, roleIdsStr)
		}

		queryStr := strings.Join(query, " AND ")

		var authAccessRule []model.AuthAccessRule
		prefix := cmf.Conf().Database.Prefix
		cmf.NewDb().Debug().Table(prefix+"auth_access access").Select("access.*,r.name").
			Joins("INNER JOIN "+prefix+"auth_rule r ON access.rule_id = r.id").Where(queryStr, queryArgs).Scan(&authAccessRule)
		session.Set("authAccessRule", authAccessRule)
		_ = session.Save()

		return authAccessRule

	}

	return res.([]model.AuthAccessRule)

}

func CurrentDate() (string,string, string){
	year,month,day := time.Now().Date()
	yearStr := strconv.Itoa(year)
	monthStr := strconv.Itoa(int(month))
	if month < 10 {
		monthStr = "0"+monthStr
	}
	dayStr := strconv.Itoa(day)
	if day < 10 {
		dayStr = "0"+dayStr
	}

	return yearStr,monthStr,dayStr
}

func SetIncr(key string) int64 {
	val ,err := cmf.NewRedisDb().Incr(key).Result()
	if err != nil {
		_,_,line,_ := runtime.Caller(0)
		fmt.Println("redis err"+strconv.Itoa(line),err.Error())
	}
	return val
}

