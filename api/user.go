package api

import (
	"database/sql"
	db "github.com/Quyen-2211/simplebank/db/sqlc"
	"github.com/Quyen-2211/simplebank/db/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" biding:"required"`
	Password string `json:"password" biding:"required,min=6"`
	Fullname string `json:"full_name" biding:"required"`
	Email    string `json:"email" biding:"required,email"`
}

type userResponse struct {
	Username         string    `json:"username"`
	Fullname         string    `json:"fullname"`
	Email            string    `json:"email"`
	PasswordChangeAt time.Time `json:"password_Change_At"`
	CreateAt         time.Time `json:"createAt"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:         user.Username,
		Fullname:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: user.PasswordChangedAt,
		CreateAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	//Nei khach hang nhap thon tin loi thi tra ve loi
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	hashedPassword, err := util.HashPassword(req.Password)

	//bay gio la tao tai khoan nguoi dung
	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.Fullname,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}

	//goi tao  tai khoan neu loi tra ve loi
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}

	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// phan dang ky
type loginUserRequest struct {
	Username string `json:"username" biding:"required"`
	Password string `json:"password" biding:"required,min=6"`
}
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//Sau khi xong thi kiem tra xem Password co dung khong

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	//chi khi password dung thi moi tao token cho nguoi dung
	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
