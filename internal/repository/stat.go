package repository

import (
	"URLProject/internal/delivery/payload"
	"URLProject/internal/stat"
	"URLProject/pkg/db"
	"gorm.io/datatypes"
	"time"
)

type StatRepository struct {
	database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{
		database: database,
	}
}

func (sr *StatRepository) AddClick(linkId uint) {
	var statistic stat.Stat
	currDate := datatypes.Date(time.Now())
	sr.database.DB.Find(&statistic, "link_id = ? and date = ?", linkId, currDate)
	if statistic.ID == 0 {
		sr.database.DB.Create(&stat.Stat{
			LinkId: linkId,
			Clicks:  1,
			Date:   currDate,
		})
	} else {
		statistic.Clicks += 1
		sr.database.DB.Save(&statistic)
	}
}

func (sr *StatRepository) GetStats(by string, from, to time.Time) []payload.GetStatResponse {
	var stats []payload.GetStatResponse
	var selectQuery string
	switch by {
	case "month":
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	case "day":
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	}
	sr.database.DB.
		Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	return stats

}