package repository

import (
	"gorm.io/gorm"
	"projectIraUw/infra/model"
	"projectIraUw/user_case/response"
)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) ListRepository {
	return ListRepository{db}
}

func (l ListRepository) ListAll() ([]response.ListAll, error) {
	var entity model.List

	filter := l.db.Model(&entity)
	resul, err := filter.Rows()
	if err != nil {
		return nil, err
	}
	defer resul.Close()

	var accountList []response.ListAll
	for resul.Next() {
		var list model.List
		resul.Scan(
			&list.ID,
			&list.Canal,
			&list.Arithmetic,
			&list.Quantidade,
			&list.Descricao,
			&list.Pessoa,
			&list.Data,
		)
		accountList = append(accountList, list.ToDomain())

	}
	return accountList, nil
}
