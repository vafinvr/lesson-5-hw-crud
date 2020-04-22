package migrations

import (
	"lesson-5-hw-crud/interfaces"
)

type mgr struct {
	fn  func() error
	idx int
}

type migrations struct {
	db interfaces.DBRepo
	fn []*mgr
}

func New(db interfaces.DBRepo) *migrations {
	m := migrations{db: db}

	m.fn = append(m.fn, &mgr{fn: m.init1, idx: 1})

	return &m
}

func (m *migrations) Up() error {
	for _, mg := range m.fn {
		if err := mg.fn(); err != nil {
			return err
		}
	}

	return nil
}
