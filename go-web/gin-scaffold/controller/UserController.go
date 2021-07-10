package controller

import (
	"fmt"
	"gin-scaffold/model"
	"gin-scaffold/service"
	"gin-scaffold/util"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserController struct {
	ApiVersion string
	Service    *service.UserService
}

type Team struct {
	Id int64
	Gk string
}

const (
	CALLBACK_REQ_NUM = 1
)

// get controller
func (uc *UserController) getCtl() *UserController {
	var svc *service.UserService
	return &UserController{"v1", svc}
}

// create user
func (uc *UserController) CreateUser(c *gin.Context) {
	var jobIdUint64 uint64
	var jobIdErrConv error

	teamId := c.PostForm("teamid")
	teamName := c.PostForm("teamname")
	resourceName := c.PostForm("resourcename")
	attackEvent := c.PostForm("attackevent")
	teamIdInt64, _ := strconv.ParseInt(teamId, 10, 64)
	atimeIdUint64 := time.Now().Unix()
	jobId, jobIdExists := c.GetPostForm("jobid")
	if jobIdExists {
		jobIdUint64, jobIdErrConv = strconv.ParseUint(jobId, 10, 64)
		if jobIdErrConv != nil {
			util.SendError(c, jobIdErrConv.Error())
			return
		}
	}

	user := &model.User{
		TeamId:       teamIdInt64,
		TeamName:     teamName,
		AttackTime:   atimeIdUint64,
		JobId:        jobIdUint64,
		AttackEvent:  attackEvent,
		ResourceName: resourceName,
		BaseModel:    model.BaseModel{},
	}

	query := map[string]interface{}{"team_id": teamId, "job_id": jobId}
	userQuery, errQuery := uc.getCtl().Service.GetUserByJobId(query)
	if errQuery == nil && userQuery.JobId == jobIdUint64 && userQuery.TeamId == teamIdInt64 {
		util.SendError(c, "user exists !")
		return
	}

	err := uc.getCtl().Service.CreateUser(user)
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
	// get user type
	userType := c.DefaultQuery("usertype", "1")
	// get users
	users, err := uc.getCtl().Service.FindAllUserByPagesWithKeys(
		map[string]interface{}{"user_type": userType},
		map[string]interface{}{},
		currentPageInt,
		pageSizeInt,
		&totalRows)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    users,
		Page: util.Pagination{
			PageSize:    pageSizeInt,
			CurrentPage: currentPageInt,
			TotalRows:   totalRows,
			TotalPages:  totalPages,
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

	keys := make(map[string]interface{})
	keyOpts := make(map[string]interface{})

	sTeamId, sTeamIdExist := c.GetQuery("steamid")
	if sTeamIdExist {
		keys["team_id"] = sTeamId
	}

	sTeamIdOpt, sTeamIdOptExist := c.GetQuery("steamidOpt")
	if sTeamIdOptExist {
		keyOpts["team_id"] = sTeamIdOpt
	}

	sTeamName, sTeamNameExist := c.GetQuery("steamname")
	if sTeamNameExist {
		keys["team_name"] = sTeamName
	}

	sTeamNameOpt, sTeamNameOptExist := c.GetQuery("steamnameOpt")
	if sTeamNameOptExist {
		keyOpts["team_name"] = sTeamNameOpt
	}

	sJobId, sJobIdExist := c.GetQuery("sjobid")
	if sJobIdExist {
		keys["job_id"] = sJobId
	}

	sJobIdOpt, sJobIdOptExist := c.GetQuery("sjobidOpt")
	if sJobIdOptExist {
		keyOpts["job_id"] = sJobIdOpt
	}

	sApp, sAppExist := c.GetQuery("sapp")
	if sAppExist {
		keys["app"] = sApp
	}

	sAppOpt, sAppOptExist := c.GetQuery("sappOpt")
	if sAppOptExist {
		keyOpts["app"] = sAppOpt
	}

	sUserType, sUserTypeExist := c.GetQuery("susertype")
	if sUserTypeExist {
		keys["user_type"] = sUserType
	}

	sUserTypeOpt, sUserTypeOptExist := c.GetQuery("susertypeOpt")
	if sUserTypeOptExist {
		keyOpts["user_type"] = sUserTypeOpt
	}

	users, err := uc.getCtl().Service.SearchUserByPagesWithKeys(keys, keyOpts, currentPageInt, pageSizeInt, &totalRows)

	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	if totalRows%int64(pageSizeInt) != 0 {
		totalPages = totalRows/int64(pageSizeInt) + 1
	} else {
		totalPages = totalRows / int64(pageSizeInt)
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    users,
		Page: util.Pagination{
			PageSize:    pageSizeInt,
			CurrentPage: currentPageInt,
			TotalRows:   totalRows,
			TotalPages:  totalPages,
		},
	})
}

// get user by id
func (uc *UserController) GetUserByID() {
}

// update user
func (uc *UserController) UpdateUser(c *gin.Context) {
	var jobIdUint64 uint64
	var jobIdErrConv error

	uid, exists := c.GetPostForm("id")
	if !exists {
		util.SendError(c, "id is null")
		return
	}

	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		util.SendError(c, "id conv failed")
		return
	}

	user, err := uc.getCtl().Service.FindUserById(uidUint64)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	teamId := c.PostForm("teamid")
	teamName := c.PostForm("teamname")
	resourceName := c.PostForm("resourcename")
	attackEvent := c.PostForm("attackevent")
	teamIdInt64, _ := strconv.ParseInt(teamId, 10, 64)
	atimeIdUint64 := time.Now().Unix()

	jobId, jobIdExists := c.GetPostForm("jobid")
	if jobIdExists {
		jobIdUint64, jobIdErrConv = strconv.ParseUint(jobId, 10, 64)
		if jobIdErrConv != nil {
			util.SendError(c, jobIdErrConv.Error())
			return
		}

		user.JobId = jobIdUint64
	}

	userType, userTypeExists := c.GetPostForm("usertype")
	if userTypeExists {
		userTypeInt8, userTypeErrConv := strconv.Atoi(userType)
		if userTypeErrConv != nil {
			util.SendError(c, jobIdErrConv.Error())
			return
		}
		user.UserType = int8(userTypeInt8)
	}

	user.TeamId = teamIdInt64
	user.TeamName = teamName
	user.AttackTime = atimeIdUint64
	user.ResourceName = resourceName
	user.AttackEvent = attackEvent

	rowsAffected, updateErr := uc.getCtl().Service.UpdateUser(uidUint64, user)
	if updateErr != nil {
		util.SendError(c, updateErr.Error())
		return
	}

	if rowsAffected == 0 {
		util.SendError(c, "数据不存在")
		return
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "update user successful",
		Data:    rowsAffected,
	})
}

// update user
func (uc *UserController) UpdateUserByUserType(c *gin.Context) {
	var jobIdErrConv error

	uid, exists := c.GetPostForm("id")
	if !exists {
		util.SendError(c, "id is null")
		return
	}

	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		util.SendError(c, "id conv failed")
		return
	}

	user, err := uc.getCtl().Service.FindUserById(uidUint64)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	userType, userTypeExists := c.GetPostForm("usertype")
	if userTypeExists {
		userTypeInt8, userTypeErrConv := strconv.Atoi(userType)
		if userTypeErrConv != nil {
			util.SendError(c, jobIdErrConv.Error())
			return
		}
		user.UserType = int8(userTypeInt8)
		// update user
		rowsAffected, updateErr := uc.getCtl().Service.UpdateUser(uidUint64, user)
		if updateErr != nil {
			util.SendError(c, updateErr.Error())
			return
		}

		if rowsAffected == 0 {
			util.SendError(c, "数据不存在")
			return
		}

		util.SendMessage(c, util.Message{
			Code:    0,
			Message: "update user successful",
			Data:    rowsAffected,
		})
	}
}

// delete user
func (uc *UserController) DeleteUser(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		util.SendError(c, "id is null")
		return
	}
	fmt.Println("uid", uid)
	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		util.SendError(c, "id conv failed")
		return
	}

	rowsAffected, delErr := uc.getCtl().Service.DeleteUser(uidUint64)
	if delErr != nil {
		util.SendError(c, delErr.Error())
		return
	}

	if rowsAffected == 0 {
		util.SendError(c, "数据不存在")
		return
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "delete user successful",
		Data:    strconv.Itoa(int(rowsAffected)) + " 条记录数受影响",
	})

}

// get all users callback-get
func (uc *UserController) GetAllUsersGetCallBack(c *gin.Context) {
	users, err := uc.getCtl().Service.FindAllUsersByCount(CALLBACK_REQ_NUM)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	// test team get request
	var teams []Team
	for _, user := range users {
		team := Team{
			Id: user.TeamId,
			Gk: user.TeamName,
		}
		teams = append(teams, team)
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    teams,
	})
}

// get all users callback-post
func (uc *UserController) GetAllUsersPostCallBack(c *gin.Context) {

	name := c.PostForm("name")
	quota := c.PostForm("quota")
	param := c.PostForm("param")

	log.Println("get all users callback-post:", "name", name, "quota", quota, "param", param)
	users, err := uc.getCtl().Service.FindAllUsersByCount(CALLBACK_REQ_NUM)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	// test team get request
	var teams []Team
	for _, user := range users {
		team := Team{
			Id: user.TeamId,
			Gk: user.TeamName,
		}
		teams = append(teams, team)
	}

	util.SendMessage(c, util.Message{
		Code:    0,
		Message: "OK",
		Data:    teams,
	})
}
