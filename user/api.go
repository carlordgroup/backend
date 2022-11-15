package user

import (
	"carlord/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service struct {
	client *ent.Client
}

type Authenticate interface {
	MustLogin() gin.HandlerFunc
	GetAccountUser() gin.HandlerFunc
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRoutes, auth Authenticate) {
	group.Use(auth.MustLogin())
	group.POST("/", s.post)
	group.Use(auth.GetAccountUser())
	group.GET("/", s.get)
}

// UserInfoGet godoc
// @Tags user
// @Summary read a user info
// @Accept json
// @Produce json
// @Success 200 {object} ent.User
// @Router /user/ [get]
func (s *service) get(ctx *gin.Context) {
	u := ctx.MustGet("account").(*ent.Account)
	ctx.JSON(http.StatusOK, u.Edges.User)
}

// UserInfoPost godoc
// @Tags user
// @Param request body ent.User true "updated info"
// @Summary update a user info
// @Accept json
// @Produce json
// @Success 200 {object} ent.User
// @Router /user/ [post]
func (s *service) post(ctx *gin.Context) {
	var u ent.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	updatedUser, err := s.client.User.UpdateOneID(ctx.MustGet("id").(int)).
		SetAddress(u.Address).
		SetBirthday(u.Birthday).
		SetDriverLicenseCountry(u.DriverLicenseCountry).
		SetDriverLicenseID(u.DriverLicenseID).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetPostalCode(u.PostalCode).
		SetTel(u.Tel).
		Save(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
