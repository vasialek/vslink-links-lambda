package interfaces

import "vslink-links-lambda/models"

type LinkRepositoryInterface interface {
	GetAllLinks() ([]models.Link, error)
}
