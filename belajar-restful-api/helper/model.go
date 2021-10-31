package helper

import (
	"belajar-golang/belajar-restful-api/model/domain"
	"belajar-golang/belajar-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponses []web.CategoryResponse
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, ToCategoryResponse(category))
	}

	return categoriesResponses
}
