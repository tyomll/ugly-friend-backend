package handler

import (
	"fmt"
	"strconv"
	"time"
	"ugly-friend/models"
	"ugly-friend/utils"
)

func (h *Handler) Register(req *models.CreateUserReq) (*models.CreateUserRes, error) {
	if req.Username == "" || req.Password == "" {
		return nil, fmt.Errorf("username_or_pass_empty")
	}

	err := h.core.Methods.CreateUser(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	jwtToken, err := utils.GenerateJWT()
	if err != nil {
		return nil, err
	}

	response := &models.CreateUserRes{Token: jwtToken, ExpiresAt: strconv.FormatInt(time.Now().Add(time.Hour*24*30).Unix(), 10)}
	return response, nil
}
