package model

import "projectIraUw/user_case/response"

type List struct {
	ID         int    `gorm:"PrimaryKey;autoIncrement"`
	Canal      string `gorm:"column:canal"`
	Arithmetic string `gorm:"column:arithmetic"`
	Quantidade string `gorm:"column:quantidade"`
	Descricao  string `gorm:"column:descricao"`
	Pessoa     string `gorm:"column:pessoa"`
	Data       string `gorm:"column:data"`
}

type ListEntity interface {
	ToDomain() *response.ListAll
}

func (ref *List) ToDomain() response.ListAll {
	return response.ListAll{
		ID:         ref.ID,
		Canal:      ref.Canal,
		Arithmetic: ref.Arithmetic,
		Quantidade: ref.Quantidade,
		Descricao:  ref.Descricao,
		Pessoa:     ref.Pessoa,
		Data:       ref.Data,
	}
}

func (ref *List) TableName() string {
	return "lista"
}
