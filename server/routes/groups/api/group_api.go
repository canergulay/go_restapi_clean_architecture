package api

import (
	"backend/server/routes/groups/model"
	"backend/server/routes/groups/service"

	"github.com/gin-gonic/gin"
)

type GroupApi struct {
	service *service.GroupService
}

func NewGroupApi(service *service.GroupService) *GroupApi {
	return &GroupApi{service: service}
}

func (a *GroupApi) QueryGroupInfo(c *gin.Context) {
	// TO BE EDITED.
}

func (a *GroupApi) CreateGroupApi(c *gin.Context) {
	var modelToCreate model.Group
	// LETS ADD A RANDOM & DEFAULT THUMBNAIL TO GROUP CREATED

	c.BindJSON(&modelToCreate)
	_, err := a.service.CreateGroup(modelToCreate)
	if err != nil {
		c.JSON(500, "error")
	}
	c.JSON(200, modelToCreate)
}

func (a *GroupApi) GetGroupsByNameAndLocaleApi(c *gin.Context) {

	var requestBody GetGroupsRequest
	c.BindJSON(&requestBody)

	resp, err := a.service.GetGroupsByNameAndLocale(requestBody.Locale, requestBody.Name)
	if err != nil {
		c.JSON(500, "error")
	}
	c.JSON(200, resp)
}
