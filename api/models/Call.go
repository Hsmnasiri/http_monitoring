package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// EndPointCalls - Object for storing endpoints call details
type EndPointCalls struct {
	gorm.Model
	EndPointID   uint64 `gorm:"index;not null"`
	RequestIP    string
	ResponseCode int
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//  func (ep Urls,) SaveCall(db *gorm.DB) EndPointCalls {
// 	epCall := EndPointCalls{
// 		EndPointID:   ep.ID,
// 		RequestIP:    
// 		ResponseCode: context.GetStatusCode(),
// 	}
// }
func (ep *EndPointCalls) Prepare() {
	ep.EndPointID = 0
	ep.RequestIP = "test"
	ep.ResponseCode =  400
	ep.CreatedAt = time.Now()
	ep.UpdatedAt = time.Now()
}

func (c *EndPointCalls) SaveCall(db *gorm.DB) (*EndPointCalls, error) {

	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &EndPointCalls{}, err
	}
	return c, nil
}