package rpc

import (
	"context"
	"fmt"
	"log"
)

// 两边都要

type GetUserByIDReq struct {
	UserId int `json:"userId"`
}

type GetUserByIDResp struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

// client

type CaseUserRPCService struct {
	GetUserByID func(context.Context, *GetUserByIDReq) (*GetUserByIDResp, error)
}

func (t *CaseUserRPCService) GetServiceName() string {
	return "user-rpc-service"
}

// server

type CaseUserService struct{}

func (t *CaseUserService) GetServiceName() string {
	return "user-rpc-service"
}

func (t *CaseUserService) GetUserByID(ctx context.Context, req *GetUserByIDReq) (*GetUserByIDResp, error) {
	log.Printf("flowId: %s", ctx.Value("flowId").(string))
	if req.UserId == 11 {
		return &GetUserByIDResp{UserID: 11, Username: "aa"}, nil
	} else if req.UserId == 22 {
		return &GetUserByIDResp{UserID: 22, Username: "bb"}, nil
	} else {
		return nil, fmt.Errorf("user id [%d] not found", req.UserId)
	}
}
