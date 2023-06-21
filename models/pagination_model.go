package models

import (
	"gorm.io/gorm"
	"math"
)

var (
	PageLimit int
	OrderBy   string
)

type Pagination struct {
	Limit       int               `json:"limit,omitempty;query:limit"`
	Page        int               `json:"page,omitempty;query:page"`
	Sort        string            `json:"sort,omitempty;query:sort"`
	TotalRows   int64             `json:"total_rows"`
	TotalPages  int               `json:"total_pages"`
	Permissions map[string]string `json:"permissions,omitempty"`
	Rows        interface{}       `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetTotalPages() {
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.GetLimit())))
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = PageLimit
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id " + OrderBy
	}
	return p.Sort
}

func Paginate(sortBy string, value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.GetTotalPages()
	pagination.Sort = sortBy

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
