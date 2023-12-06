package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GeovaneCavalcante/api-users/internal/dto"
	"github.com/GeovaneCavalcante/api-users/internal/entity"
	"github.com/GeovaneCavalcante/api-users/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        db,
		Jwt:           jwt,
		JwtExperiesIn: jwtExperiesIn,
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.GetJWTInput true "user request"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users/generate_token [post]
func (p *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := p.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, err := p.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Hour * time.Duration(p.JwtExperiesIn)).Unix(),
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	acessToken := dto.GetJWTOutput{
		AcessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}

// Create user godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
func (p *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errUser := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errUser)
		return
	}
	err = p.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errUser := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errUser)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
