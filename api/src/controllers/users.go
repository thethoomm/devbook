package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("signup"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	respository := repositories.NewUserRepository(db)
	users, err := respository.Find(nameOrUsername)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func FindUniqueUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.FindById(ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if ID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário se não o seu"))
		return
	}

	fmt.Println(userIDToken)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("edit"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	respository := repositories.NewUserRepository(db)
	if err := respository.Update(ID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if ID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário se não o seu"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Delete(ID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Follow(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Unfollow(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	followers, err := repository.FindFollowers(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.FindFollowing(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível alterar a senha de um usuário que não seja o seu"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var password models.Password

	if err := json.Unmarshal(body, &password); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	passwordStoraged, err := repository.FindPassword(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(passwordStoraged, password.Current); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	passwordHashed, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.UpdatePassword(userID, string(passwordHashed)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
