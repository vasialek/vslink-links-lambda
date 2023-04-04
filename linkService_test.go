package services

import (
	"testing"
)

func TestX(t *testing.T) {
	t.Error("Fake")
}

// func TestGetAllLinks(t *testing.T) {
// 	sut := services.NewLinkService(NewMockLinkRepository())

// 	// sut := services.NewLinkService(NewMockLinkRepository([]models.Link))

// 	actual, err := sut.GetAllLinks()

// 	if err != nil {
// 		t.Error("expecting no error")
// 	}
// 	if len(actual) != 1 {
// 		t.Errorf("Expecting to return 1 link, but got %d\n", len(actual))
// 	}
// }

// type MockLinkRepository struct {
// 	links []models.Link
// }

// func NewMockLinkRepository() MockLinkRepository {
// 	return MockLinkRepository{}
// }

// // func NewMockLinkRepository(links []models.Link) *MockLinkRepository {
// // 	return &MockLinkRepository{
// // 		links: links,
// // 	}
// // }

// func (mlr *MockLinkRepository) GetAllLinks() ([]models.Link, error) {
// 	return mlr.links, nil
// }
