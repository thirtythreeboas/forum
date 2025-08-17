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
		if err == cerr.ErrUserDoesntExist {
			pkg.WriteError(w, http.StatusNotFound, cerr.ErrCantFindUser.Error())
			return
		}

		if err == cerr.ErrForumAlreadyExists {
			pkg.WriteResponse(w, forum, "", http.StatusConflict)
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, forum, "", http.StatusCreated)
}

func (h *Handler) GetForumInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrWrongSlugProvided.Error())
		return
	}

	existingForum, err := h.service.GetForumInfo(slug)

	if err != nil {
		if err == cerr.ErrForumDoesntExist {
			pkg.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, existingForum, "", http.StatusOK)
}

func (h *Handler) CreateThread(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrWrongSlugProvided.Error())
		return
	}

	newThread := &model.NewThread{}
	newThread.Slug = slug

	err = json.Unmarshal(body, newThread)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	thread, err := h.service.CreateThread(newThread)
	if err != nil {
		if err == cerr.ErrUserDoesntExist || err == cerr.ErrForumDoesntExist {
			pkg.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		if err == cerr.ErrThreadAlreadyExists {
			pkg.WriteResponse(w, thread, "", http.StatusConflict)
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, thread, "", http.StatusOK)
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
	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrNicknameParamNotProvided.Error())
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

func (h *Handler) GetThreadInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrWrongSlugProvided.Error())
		return
	}

	thread, err := h.service.GetThreadInfo(slug)
	if err != nil {
		if err == cerr.ErrThreadDoesntExist {
			pkg.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteResponse(w, thread, "", http.StatusOK)
}

// func (h *Handler) CreatePosts(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		pkg.WriteError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	slug, ok := vars["slug"]
// 	if !ok {
// 		pkg.WriteError(w, http.StatusBadRequest, cerr.ErrWrongSlugProvided.Error())
// 		return
// 	}

// 	var newPosts []*model.NewPost

// 	err = json.Unmarshal(body, newPosts)
// 	if err != nil {
// 		pkg.WriteError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	threads, err := h.service.CreateThreads(slug, newPosts)
// }
