package services

import (
	"vslink-links-lambda/interfaces"
	"vslink-links-lambda/models"
)

type LinkService struct {
	linkRepository *interfaces.LinkRepositoryInterface
}

func NewLinkService(linkRepository interfaces.LinkRepositoryInterface) *LinkService {
	return &LinkService{
		linkRepository: &linkRepository,
	}
}

func (ls *LinkService) GetAllLinks() ([]models.Link, error) {
	panic("asdfsdf")
}
