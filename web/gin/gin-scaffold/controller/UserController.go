package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin-scaffold/middleware"
	"gin-scaffold/model"
	"gin-scaffold/service"
	"gin-scaffold/util"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// controller struct
type UserController struct {
	apiVersion string
	Service    *service.UserService
}

// get controller
func (uc *UserController) getCtl() *UserController {
	var svc *service.UserService
	return &UserController{"v1", svc}
}

// login
func (uc *UserController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "用户登录",
	})
}

// do login
func (uc *UserController) DoLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var us *service.UserService
	user, errFind := us.FindUserByEmail(email)

	if user == nil || errFind != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "user not find or find err !",
		})
	} else {
		errPasswd := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if user.Email == email && errPasswd == nil {
			middleware.SaveAuthSession(c, user.ID)
			uc.UserHome(c)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"error": "email or password error !",
			})
		}
	}
}

// user home
func (uc *UserController) UserHome(c *gin.Context) {
	var us *service.UserService

	var currentPageInt, pageSizeInt = util.CURRENT_PAGE, util.PAGE_SIZE
	var totalRows, totalPages int64
	pageSizeInt = viper.GetInt("PAGE_SIZE")
	currentPage, cpExist := c.GetQuery("currentpage")
	if cpExist {
		currentPageInt, _ = strconv.Atoi(currentPage)
	}

	pageSize, psExist := c.GetQuery("pagesize")
	if psExist {
		pageSizeInt, _ = strconv.Atoi(pageSize)
	}

	users, errUsers := us.FindAllUserByPagesWithKeys(
		map[string]interface{}{},
		map[string]interface{}{},
		currentPageInt,
		pageSizeInt,
		&totalRows)

	if errUsers != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errUsers,
		})
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "管理后台",
		"users": users,
		"pages": util.Pagination{
			PageSize:        pageSizeInt,
			CurrentPage:     currentPageInt,
			TotalRows:       totalRows,
			TotalPages:      totalPages,
			PreCurrentPage:  currentPageInt - 1,
			NextCurrentPage: currentPageInt + 1,
		},
	})
}

// logout
func (uc *UserController) Logout(c *gin.Context) {
	middleware.ClearAuthSession(c)
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "用户登录",
	})
}

// get all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var currentPageInt, pageSizeInt = util.CURRENT_PAGE, util.PAGE_SIZE
	var totalRows, totalPages int64
	pageSizeInt = viper.GetInt("PAGE_SIZE")
	currentPage, cpExist := c.GetQuery("currentpage")
	if cpExist {
		currentPageInt, _ = strconv.Atoi(currentPage)
	}

	pageSize, psExist := c.GetQuery("pagesize")
	if psExist {
		pageSizeInt, _ = strconv.Atoi(pageSize)
	}

	// data option setting
	dataOrder, dataOrderExist := c.GetQuery("dataOrder")
	if !dataOrderExist {
		dataOrder = "id desc"
	}

	dataSelect, dataSelectExist := c.GetQuery("dataSelect")
	if !dataSelectExist {
		dataSelect = ""
	}

	dataWhereMap := map[string]interface{}{}
	dataWhere, dataWhereExist := c.GetQuery("dataWhere")
	if dataWhereExist {
		err := json.Unmarshal([]byte(dataWhere), &dataWhereMap)
		if err != nil {
			util.SendError(c, err.Error())
			return
		}
	}

	dataLimitInt := 0
	dataLimit, dataLimitExist := c.GetQuery("dataLimit")
	if dataLimitExist {
		dataLimitInt, _ = strconv.Atoi(dataLimit)
	}

	daoOpt := model.DAOOption{
		Select: dataSelect,
		Order:  dataOrder,
		Where:  dataWhereMap, //map[string]interface{}{},
		Limit:  dataLimitInt,
	}

	users, err := uc.getCtl().Service.FindAllUserByPagesWithKeys(
		map[string]interface{}{},
		map[string]interface{}{},
		currentPageInt,
		pageSizeInt,
		&totalRows,
		daoOpt)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	c.HTML(http.StatusOK, "users.tmpl", gin.H{
		"title": "用户管理",
		"users": users,
		"pages": util.Pagination{
			PageSize:        pageSizeInt,
			CurrentPage:     currentPageInt,
			TotalRows:       totalRows,
			TotalPages:      totalPages,
			PreCurrentPage:  currentPageInt - 1,
			NextCurrentPage: currentPageInt + 1,
		},
	})
}

// search users by pages with keys
func (uc *UserController) SearchUsersByKeys(c *gin.Context) {
	var currentPageInt, pageSizeInt = util.CURRENT_PAGE, util.PAGE_SIZE
	var totalRows, totalPages int64
	pageSizeInt = viper.GetInt("PAGE_SIZE")
	currentPage, cpExist := c.GetQuery("currentpage")
	if cpExist {
		currentPageInt, _ = strconv.Atoi(currentPage)
	}

	pageSize, psExist := c.GetQuery("pagesize")
	if psExist {
		pageSizeInt, _ = strconv.Atoi(pageSize)
	}

	//search option
	keys := make(map[string]interface{})
	keyOpts := make(map[string]interface{})
	//get search key
	searchKey, searchKeyExist := c.GetQuery("searchKey")
	if searchKeyExist {
		keys["name"] = searchKey
		keys["email"] = searchKey
		keys["age"] = searchKey
		keys["role"] = searchKey
		log.Println("000000000", keys)
	} else {
		//search key
		name, nameExist := c.GetQuery("name")
		if nameExist {
			keys["name"] = name
		}

		email, emailExist := c.GetQuery("email")
		if emailExist {
			keys["email"] = email
		}

		age, ageExist := c.GetQuery("age")
		if ageExist {
			keys["age"] = age
		}

		role, roleExist := c.GetQuery("role")
		if roleExist {
			keys["role"] = role
		}
	}

	//search value options
	searchKeyOpt, searchKeyOptExist := c.GetQuery("searchKeyOpt")
	if searchKeyOptExist {
		keyOpts["searchKey"] = searchKeyOpt
	}

	nameOpt, nameOptExist := c.GetQuery("nameOpt")
	if nameOptExist {
		keyOpts["name"] = nameOpt
	}

	emailOpt, emailOptExist := c.GetQuery("emailOpt")
	if emailOptExist {
		keyOpts["email"] = emailOpt
	}

	ageOpt, ageOptExist := c.GetQuery("ageOpt")
	if ageOptExist {
		keyOpts["age"] = ageOpt
	}

	roleOpt, roleOptExist := c.GetQuery("roleOpt")
	if roleOptExist {
		keyOpts["role"] = roleOpt
	}

	// data option setting
	dataOrder, dataOrderExist := c.GetQuery("dataOrder")
	if !dataOrderExist {
		dataOrder = "id desc"
	}

	dataSelect, dataSelectExist := c.GetQuery("dataSelect")
	if !dataSelectExist {
		dataSelect = ""
	}

	dataWhereMap := map[string]interface{}{}
	dataWhere, dataWhereExist := c.GetQuery("dataWhere")
	if dataWhereExist {
		err := json.Unmarshal([]byte(dataWhere), &dataWhereMap)
		if err != nil {
			util.SendError(c, err.Error())
			return
		}
	}

	dataLimitInt := 0
	dataLimit, dataLimitExist := c.GetQuery("dataLimit")
	if dataLimitExist {
		dataLimitInt, _ = strconv.Atoi(dataLimit)
	}

	daoOpt := model.DAOOption{
		Select: dataSelect,
		Order:  dataOrder,
		Where:  dataWhereMap, //map[string]interface{}{},
		Limit:  dataLimitInt,
	}

	users, err := uc.getCtl().Service.SearchUserByPagesWithKeys(keys,
		keyOpts,
		currentPageInt,
		pageSizeInt,
		&totalRows,
		daoOpt)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	c.HTML(http.StatusOK, "users.tmpl", gin.H{
		"title": "用户管理",
		"users": users,
		"pages": util.Pagination{
			PageSize:        pageSizeInt,
			CurrentPage:     currentPageInt,
			TotalRows:       totalRows,
			TotalPages:      totalPages,
			PreCurrentPage:  currentPageInt - 1,
			NextCurrentPage: currentPageInt + 1,
		},
	})
}

// get user by id
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, exists := c.Params.Get("id")
	if !exists {
		util.SendError(c, "id is null")
		return
	}

	idUint64, errConv := strconv.ParseUint(id, 10, 64)
	if errConv != nil {
		util.SendError(c, "id conv failed")
		return
	}

	user, err := uc.getCtl().Service.FindUserById(idUint64)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    user,
	})
}

// add user tmpl
func (uc *UserController) AddUser(c *gin.Context) {
	c.HTML(http.StatusOK, "user_add.tmpl", gin.H{
		"title": "添加用户",
	})
}

// create user
func (uc *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password, exists := c.GetPostForm("password")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "password is null",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "bcrypt password err",
		})
	}

	password = string(hash)
	email := c.PostForm("email")
	role := c.PostForm("role")
	age, _ := strconv.Atoi(c.PostForm("age"))
	user := &model.User{
		Name:         name,
		Password:     password,
		Email:        email,
		Role:         role,
		Age:          age,
		Birthday:     time.Now(),
		MemberNumber: sql.NullString{},
		BaseModel:    model.BaseModel{},
	}

	errCreate := uc.getCtl().Service.CreateUser(user)
	if errCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	//redirect
	c.Redirect(http.StatusMovedPermanently, "/users")
}

// update user
func (uc *UserController) UpdateUser(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}

	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		util.SendError(c, "get uid error !")
		return
	}

	user, err := uc.getCtl().Service.FindUserById(uidUint64)
	if err != nil {
		panic(" get user error !")
	}

	//update data
	updateDataEnabled, updateDataExist := c.GetPostForm("updatedata")
	if updateDataExist && updateDataEnabled == "true" {
		c.HTML(http.StatusOK, "user_update.tmpl", gin.H{
			"title": "更新用户",
			"user":  user,
		})
		return
	}

	user.ID = uidUint64
	name, nameExist := c.GetPostForm("name")
	if nameExist {
		user.Name = name
	}

	email, emailExist := c.GetPostForm("email")
	if emailExist {
		user.Email = email
	}

	age, ageExist := c.GetPostForm("age")
	if ageExist {
		ageInt, _ := strconv.Atoi(age)
		user.Age = ageInt
	}

	role, roleExist := c.GetPostForm("role")
	if roleExist {
		user.Role = role
	}
	//update user
	rowsAffected, updateErr := uc.getCtl().Service.UpdateUser(uidUint64, user)
	if updateErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": updateErr,
		})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": rowsAffected,
	// })

	log.Println("rows affected: ", rowsAffected)
	//redirect
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}
	fmt.Println("uid", uid)
	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	rowsAffected, delErr := uc.getCtl().Service.DeleteUser(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete user error",
		})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": rowsAffected,
	// })

	log.Println("rows affected: ", rowsAffected)
	//redirect
	c.Redirect(http.StatusMovedPermanently, "/users")
}
