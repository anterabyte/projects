package gapi

import (
	"fmt"

	db "gitlab.com/d3vus/project/go/bankServer/db/sqlc"
	"gitlab.com/d3vus/project/go/bankServer/pb"
	"gitlab.com/d3vus/project/go/bankServer/token"
	"gitlab.com/d3vus/project/go/bankServer/util"
	"gitlab.com/d3vus/project/go/bankServer/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
