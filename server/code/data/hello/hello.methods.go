package hello

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/sebastianaldi17/dockerfullstack/server/code/entity"
)

func (d *Data) LogHello() error {
	_, err := d.db.Query(`
		INSERT INTO hello(time)
		VALUES ($1)
	`, time.Now())

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Data) GetLogs(req entity.HelloLogsRequest) (entity.HelloLogsResponse, error) {
	rows, err := d.db.Query(`
	SELECT id, time FROM hello
	WHERE time between $1 and $2
	LIMIT $3 OFFSET $4
	`, req.Filter.StartDate, req.Filter.EndDate, req.Pager.Limit, req.Pager.Next)

	if err != nil {
		log.Println(err)
		return entity.HelloLogsResponse{
			Metadata: entity.Metadata{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while trying to query data.",
			},
		}, err
	}

	var (
		id        int64
		timestamp time.Time
	)
	logs := make([]entity.Log, 0)
	for rows.Next() {
		err := rows.Scan(&id, &timestamp)
		if err != nil {
			log.Println(err)
			return entity.HelloLogsResponse{
				Metadata: entity.Metadata{
					StatusCode: http.StatusInternalServerError,
					Message:    "Error while trying to read rows.",
				},
			}, err
		}
		logs = append(logs, entity.Log{
			ID:        id,
			Timestamp: timestamp,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return entity.HelloLogsResponse{
			Logs: []entity.Log{},
			Metadata: entity.Metadata{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while trying to read rows.",
			},
		}, err
	}
	return entity.HelloLogsResponse{
		Logs: logs,
		Metadata: entity.Metadata{
			StatusCode: http.StatusOK,
		},
	}, nil
}
