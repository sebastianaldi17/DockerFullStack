package hello

import (
	"log"
	"net/http"
	"time"

	"github.com/sebastianaldi17/dockerfullstack/server/code/entity"
)

func (b *Business) Hello() error {
	err := b.helloData.LogHello()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (b *Business) GetLogs(req entity.HelloLogsRequest) (entity.HelloLogsResponse, error) {
	if req.Pager.Limit <= 0 {
		req.Pager.Limit = 10
	}

	if req.Filter.EndDate.IsZero() {
		req.Filter.EndDate = time.Now()
	}

	if req.Filter.StartDate.After(req.Filter.EndDate) {
		return entity.HelloLogsResponse{
			Metadata: entity.Metadata{
				StatusCode: http.StatusBadRequest,
				Message:    "Start time cannot be after end time.",
			},
		}, nil
	}

	logs, err := b.helloData.GetLogs(req)

	return logs, err
}
