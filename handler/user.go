package handler

import (
	"ugly-friend/core"
	"ugly-friend/models"
)

func NewHandler(core *core.UglyFriendCore) *Handler {
	return &Handler{core: core}
}

type Handler struct {
	core *core.UglyFriendCore
}

func (h *Handler) Register(*models.CreateUserReq) (*models.CreateUserRes, error) {

	return &models.CreateUserRes{}, nil
}
