package database

import (
	"fmt"
	"strconv"
)

func (b Book) Print() string {
	pubDate := "Unknown"
	if b.PublicationDate.Valid {
		pubDate = b.PublicationDate.String
	}

	nPages := "Unknown"
	if b.NumberOfPages.Valid {
		nPages = strconv.FormatInt(b.NumberOfPages.Int64, 10)
	}

	return fmt.Sprintf(`%s
	ISBN: %s
	Publication Date: %s
	Pages: %s`,
		b.Title, strconv.FormatInt(b.Isbn, 10), pubDate, nPages)
}
