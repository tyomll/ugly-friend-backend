package handler

import (
	"ugly-friend/core"
)

func NewHandler(core *core.UglyFriendCore) *Handler {
	return &Handler{core: core}
}

type Handler struct {
	core *core.UglyFriendCore
}
