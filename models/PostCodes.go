package models

import (
	"gorm.io/datatypes"
	"time"
)

type PostCode struct {
	ID          uint64         `json:"id" yaml:"-" gorm:"primaryKey;comment:ID"`
	Code        string         `json:"code" yaml:"code" gorm:"index:post_code_idx,sort:desc;comment:Posta Kodu"`
	City        string         `json:"city" yaml:"city" gorm:"index:city_county_town_idx,sort:desc;comment:Şehir"`
	County      string         `json:"county" yaml:"county" gorm:"index:city_county_town_idx,sort:desc;comment:İlçe"`
	Town        string         `json:"town" yaml:"town" gorm:"index:city_county_town_idx,sort:desc;comment:Semt, Bucak, Belde"`
	District    string         `json:"district" yaml:"district" gorm:"comment:Mahalle"`
	CountryCode string         `json:"country_code" yaml:"countryCode" gorm:"default:Turkiye;comment:Ülke Kodu"`
	Location    datatypes.JSON `json:"location,omitempty" yaml:"location"`
	UploadDate  time.Time      `json:"upload_date" yaml:"-" gorm:"type:date;comment:Yükleme Tarihi"`
	CreatedAt   time.Time      `json:"created_at" yaml:"-" gorm:"comment:Oluşturulma Tarihi"`
	UpdatedAt   time.Time      `json:"updated_at" yaml:"-" gorm:"comment:Güncellenme Tarihi"`
}

type Location struct {
	X, Y float64
}

type City struct {
	Name string `json:"city" yaml:"name"`
}
type County struct {
	City   string `json:"city"   yaml:"city"`
	County string `json:"county" yaml:"county"`
}
type Town struct {
	City   string `json:"city"   yaml:"city"`
	County string `json:"county" yaml:"county"`
	Town   string `json:"town"   yaml:"town"`
}
type District struct {
	Code     string `json:"code"     yaml:"code"`
	City     string `json:"city"     yaml:"city"`
	County   string `json:"county"   yaml:"county"`
	Town     string `json:"town"     yaml:"town" `
	District string `json:"district" yaml:"district"`
}
