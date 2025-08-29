package database

import (
	"fmt"
	"strconv"
)

func (b Book) Print() string {
	pubDate := "Unkown"
	if b.PublicationDate.Valid {
		pubDate = b.PublicationDate.String
	}

	nPages := "Unkown"
	if b.NumberOfPages.Valid {
		nPages = strconv.FormatInt(b.NumberOfPages.Int64, 10)
	}

	return fmt.Sprintf(`%s
ISBN: %s
Publication Date: %s
Pages: %s`,
		b.Title, b.ID, pubDate, nPages)
}
