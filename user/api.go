package user

import (
	"carlord/data"
	"carlord/ent"
	"carlord/ent/account"
	"carlord/ent/user"
	"carlord/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service struct {
	client *ent.Client
}

type Authenticate interface {
	MustLogin() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
}

func New(client *ent.Client) *service {
	s := service{
		client: client,
	}
	return &s
}

func (s *service) RegisterRouter(group gin.IRoutes, auth Authenticate) {
	group.Use(auth.MustLogin())
	group.POST("/", web.W(s.post))
	group.Use(auth.GetAccount())
	group.GET("/", web.W(s.get))
}

// UserInfoGet godoc
// @Tags user
// @Summary read a user info
// @Accept json
// @Produce json
// @Success 200 {object} ent.User
// @Router /user/ [get]
func (s *service) get(ctx *gin.Context) (int, any) {
	u := ctx.MustGet("account").(*data.Account)
	return http.StatusOK, u.Edges.User
}

// UserInfoPost godoc
// @Tags user
// @Param request body ent.User true "updated info"
// @Summary update a user info
// @Accept json
// @Produce json
// @Success 200 {object} ent.User
// @Router /user/ [post]
func (s *service) post(ctx *gin.Context) (int, any) {
	var u ent.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		return http.StatusBadRequest, err
	}
	accountID := ctx.MustGet("id").(int)
	var userEdit *ent.User
	hasUser := s.client.User.Query().Where(user.HasAccountWith(account.ID(accountID))).ExistX(ctx)
	if hasUser {
		userEdit = s.client.User.Query().Where(user.HasAccountWith(account.ID(accountID))).OnlyX(ctx)
	} else {
		userEdit = s.client.User.Create().SetAccountID(accountID).SaveX(ctx)
	}

	updatedUser, err := userEdit.Update().
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
		return http.StatusBadRequest, err
	}

	return http.StatusOK, updatedUser
}
