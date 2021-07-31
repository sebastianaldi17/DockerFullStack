package hello

import "github.com/sebastianaldi17/dockerfullstack/server/code/entity"

//go:generate mockgen -build_flags=-mod=mod -source=hello.dependencies.go -package=hello -destination=hello.dependencies.mock_test.go
type helloBusiness interface {
	Hello() error
	GetLogs(entity.HelloLogsRequest) (entity.HelloLogsResponse, error)
}
