package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Urls struct {
	ID        uint32   `gorm:"primary_key;auto_increment" json:"id"`
	Name, URL string
	Type      string          `gorm:"DEFAULT:'GET'"`
	Calls     []EndPointCalls `gorm:"ForeignKey:EndPointID"`
	Owner    User      `json:"owner"`
	OwnerID  uint32   `gorm:"not null" json:"owner_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Urls) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.URL = html.EscapeString(strings.TrimSpace(p.URL))
	p.Owner = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Urls) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}

	if p.URL == "" {
		return errors.New("Required URL")
	}

	if p.Type == "" {
		return errors.New("Required Type")
	}

	if p.OwnerID < 1 {
		return errors.New("Required owner")
	}
	return nil
}

func (url *Urls) SaveUrl(db *gorm.DB) (*Urls, error) {
	var err error
	err = db.Debug().Model(&Urls{}).Create(&url).Error
	if err != nil {
		return &Urls{}, err
	}
	if url.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", url.OwnerID).Take(&url.Owner).Error
		if err != nil {
			return &Urls{}, err
		}
	}
	return url, nil
}

func (p *Urls) FindAllUrlses(db *gorm.DB) (*[]Urls, error) {
	var err error
	Urlses := []Urls{}
	err = db.Debug().Model(&Urls{}).Limit(100).Find(&Urlses).Error
	if err != nil {
		return &[]Urls{}, err
	}
	if len(Urlses) > 0 {
		for i := range Urlses {
			err := db.Debug().Model(&User{}).Where("id = ?", Urlses[i].OwnerID).Take(&Urlses[i].Owner).Error
			if err != nil {
				return &[]Urls{}, err
			}
		}
	}
	return &Urlses, nil
}

func (url *Urls) FindUrlByID(db *gorm.DB, pid uint64) (*Urls, error) {
	var err error
	err = db.Debug().Model(&Urls{}).Where("id = ?", pid).Take(&url).Error
	if err != nil {
		return &Urls{}, err
	}
	if url.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", url.OwnerID).Take(&url.Owner).Error
		if err != nil {
			return &Urls{}, err
		}
	}
	return url, nil
}

func (p *Urls) UpdateAUrl(db *gorm.DB) (*Urls, error) {

	var err error

	err = db.Debug().Model(&Urls{}).Where("id = ?", p.ID).Updates(Urls{Name: p.Name, URL: p.URL, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Urls{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.OwnerID).Take(&p.Owner).Error
		if err != nil {
			return &Urls{}, err
		}
	}
	return p, nil
}

func (p *Urls) DeleteAUrl(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Urls{}).Where("id = ? and owner_id = ?", pid, uid).Take(&Urls{}).Delete(&Urls{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Urls not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}