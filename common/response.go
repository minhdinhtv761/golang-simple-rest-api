package common

type QueryOptions interface {
	IsEmpty() bool
}

type successResponse struct {
	Data   interface{}  `json:"data"`
	Paging QueryOptions `json:"paging,omitempty"`
	Filter QueryOptions `json:"filter,omitempty"`
	Sort   QueryOptions `json:"sort,omitempty"`
}

func NewSuccessResponse(data interface{}, paging, filter, sort QueryOptions) *successResponse {
	if paging != nil && paging.IsEmpty() {
		paging = nil
	}

	if filter != nil && filter.IsEmpty() {
		filter = nil
	}

	if sort != nil && sort.IsEmpty() {
		sort = nil
	}

	return &successResponse{data, paging, filter, sort}
}

func NewSimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil, nil, nil)
}

func NewDoneSuccessResponse() *successResponse {
	return NewSuccessResponse(
		map[string]interface{}{"success": true},
		nil,
		nil,
		nil,
	)
}
