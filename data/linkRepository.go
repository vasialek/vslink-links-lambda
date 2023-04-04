package data

import "vslink-links-lambda/models"

type LinkRepository struct {
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{}
}

// GetAllLinks returns all links in database
func (lr *LinkRepository) GetAllLinks() ([]models.Link, error) {
	return []models.Link{
		{
			LinkID:         "3a01e0b71c514c7a9da53c095b4433a1",
			LinkCategoryID: "38f4f9d5f12a416193a6c22a56d9c142",
			UserID:         "99d45738976d4a64acee1356e49044f4",
			Title:          "A-baza new messages",
			URL:            "http://www.wrk.ru/forums/search.php?action=show_new",
			Rating:         1,
		},
		{
			LinkID:         "77f83c82b7aa46debdf6c7feea95c2b3",
			LinkCategoryID: "6ea6ac0db8114f17ba717145e0990ce7",
			UserID:         "99d45738976d4a64acee1356e49044f4",
			Title:          "Turn off LED with Hall effect sensor?",
			URL:            "https://electronics.stackexchange.com/questions/137774/turn-off-led-with-hall-effect-sensor",
			Rating:         0,
		},
		{
			LinkID:         "74270b454331414a953f809bcb63c01a",
			LinkCategoryID: "74bbe5264219493688c75f4cd90b2cf4",
			UserID:         "99d45738976d4a64acee1356e49044f4",
			Title:          "My friend complaining...",
			URL:            "https://ic.pics.livejournal.com/andruxa_baun/13659942/5948712/5948712_original.jpg",
			Rating:         2,
		},
	}, nil
}

func (lr *LinkRepository) CreateLink(link models.Link) (*models.Link, error) {
	panic("Not implemented CreateLink")
}
