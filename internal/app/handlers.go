package app

import (
	"encoding/json"
	"forum/internal/model"
	"forum/pkg"
	cerr "forum/pkg/errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateForum(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	newForum := &model.ForumCreate{}

	err = json.Unmarshal(body, newForum)

	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	forum, err := h.service.CreateForum(newForum)

	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, forum, "", http.StatusCreated)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrNicknameParamNotProvided.Error())
	}

	newUser := &model.NewProfile{}
	newUser.Nickname = nickname

	err = json.Unmarshal(body, newUser)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.CreateUser(newUser)

	if err != nil {
		if err == cerr.ErrUserAlreadyExists {
			pkg.WriteResponse(w, []*model.User{user}, "", http.StatusConflict)
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, user, "", http.StatusCreated)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
	}

	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrNicknameParamNotProvided.Error())
	}

	user := &model.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	extractedUser, err := h.service.GetUser(nickname)

	if err != nil {
		if err == cerr.ErrCantFindUser {
			pkg.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())

		return
	}

	pkg.WriteResponse(w, extractedUser, "", http.StatusOK)
}

func (h *Handler) ChangeUserProfile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)

	user := &model.User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	nickname, ok := vars["nickname"]
	user.Nickname = nickname

	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrNicknameParamNotProvided.Error())
		return
	}

	extractedUser, err := h.service.ChangeProfile(user)

	if err != nil {
		if err == cerr.ErrEmailIsInUse {
			pkg.WriteError(w, http.StatusConflict, cerr.ErrCantFindUser.Error())
			return
		}
		if err == cerr.ErrUserDoesntExist {
			pkg.WriteError(w, http.StatusNotFound, cerr.ErrCantFindUser.Error())
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, extractedUser, "", http.StatusOK)
}
