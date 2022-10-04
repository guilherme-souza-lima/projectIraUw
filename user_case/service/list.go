package service

import "projectIraUw/user_case/response"

type RepositoryListINTER interface {
	ListAll() ([]response.ListAll, error)
}

type ListService struct {
	RepositoryListINTER RepositoryListINTER
}

func NewListService(RepositoryListINTER RepositoryListINTER) ListService {
	return ListService{RepositoryListINTER}
}

func (l ListService) ListAllService() ([]response.ListAll, error) {
	return l.RepositoryListINTER.ListAll()
}
