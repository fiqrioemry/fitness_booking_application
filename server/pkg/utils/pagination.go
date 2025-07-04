package utils

import "server/internal/dto"

func Paginate(total int64, page, limit int) *dto.PaginationResponse {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &dto.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}
}
