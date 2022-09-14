package fnet

func PageOffset(page, size int) int {
	return (page - 1) * size
}

func NumberOfPages(records, size int) int {
	pages := records / size
	rem := records % size

	if rem > 0 {
		return pages + 1
	} else {
		return pages
	}
}
