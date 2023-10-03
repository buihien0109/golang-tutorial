package database

import "book-app/model"

var BookList = []model.Book{
	{
		ID:          1,
		Title:       "The Hitchhiker's Guide to the Galaxy",
		Description: "The Hitchhiker's Guide to the Galaxy is a comic science fiction series created by Douglas Adams that has become popular among fans of the genre and members of the scientific community.",
		Author:      "Nguyen Van A",
	},
	{
		ID:          2,
		Title:       "The Restaurant at the End of the Universe",
		Description: "The Restaurant at the End of the Universe is the second book in the Hitchhiker's Guide to the Galaxy comedy science fiction trilogy by Douglas Adams, and is a sequel.",
		Author:      "Tran Nguyen D",
	},
	{
		ID:          3,
		Title:       "Life, the Universe and Everything",
		Description: "Life, the Universe and Everything is the third book in the five-volume Hitchhiker's Guide to the Galaxy science fiction trilogy by British writer Douglas Adams.",
		Author:      "Nguyen Van B",
	},
	{
		ID:          4,
		Title:       "So Long, and Thanks for All the Fish",
		Description: "So Long, and Thanks for All the Fish is the fourth book of the Hitchhiker's Guide to the Galaxy trilogy written by Douglas Adams. Its title is the message left by the dolphins when they departed Planet Earth just before it was demolished to make way for a hyperspace bypass, as described in The Hitchhiker's Guide to the Galaxy.",
		Author:      "Ngo Van C",
	},
}
