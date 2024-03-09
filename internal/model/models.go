package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Figures FigureModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Figures: FigureModel{DB: db},
	}
}
