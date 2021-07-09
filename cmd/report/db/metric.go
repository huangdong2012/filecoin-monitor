package db

import (
	"database/sql"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
	"time"
)

func InsertMetrics(ms []*model.Metric) error {
	if len(ms) == 0 {
		return nil
	}

	entities := make([]*Metric, 0, 0)
	for _, m := range ms {
		entity := &Metric{
			Name:   m.Name,
			Labels: utils.ToJsonWithoutError(m.Labels),
			Value:  m.Value,
			Description: sql.NullString{
				String: m.Desc,
				Valid:  true,
			},
			MetricTime: time.Unix(m.Time, 0),
			CreateTime: time.Now(),
		}
		entities = append(entities, entity)
	}

	return db.Create(&entities).Error
}
