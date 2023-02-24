package common

import "reflect"

type Paging struct {
	Page  int   `json:"page,omitempty" form:"page"`
	Limit int   `json:"limit,omitempty" form:"limit"`
	Total int64 `json:"total,omitempty"`
	// Support cursor with UID
	FakeCursor string `json:"cursor,omitempty" form:"cursor"`
	NextCursor string `json:"next_cursor,omitempty"`
}

func (page *Paging) Fulfill() {
	if page.Page <= 0 {
		page.Page = 1
	}

	if page.Limit <= 0 {
		page.Limit = 50
	}
}

func (page *Paging) IsEmpty() bool {
	return reflect.ValueOf(*page).IsZero()
}
