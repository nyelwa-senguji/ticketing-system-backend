package endpoint

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type (
	CreateCategoryResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ListCategoriesResponse struct {
		Status     int           `json:"status"`
		Message    string        `json:"message"`
		Categories []db.Category `json:"categories"`
	}

	GetCategoryRequest struct {
		Id int32 `json:"id"`
	}

	GetCategoryResponse struct {
		Status   int         `json:"status"`
		Message  string      `json:"message"`
		Category db.Category `json:"category"`
	}
)

func makeCreateCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreateCategoryRequest)

		if reflect.DeepEqual(req.CategoryName, "") {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: "Category name cannot be empty", Success: false}, nil
		}

		if reflect.DeepEqual(req.Status, "") {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: "Category status cannot be empty", Success: false}, nil
		}

		ok, err := s.CreateCategory(ctx, req)
		if err != nil {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: err.Error(), Success: false}, nil
		}

		return CreateCategoryResponse{Status: utils.StatusOK, Message: ok, Success: true}, err
	}
}

func makeGetCategoryEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCategoryRequest)
		ok, err := s.GetCategory(ctx, req.Id)
		if err != nil {
			return GetCategoryResponse{Status: utils.StatusBadRequest, Message: err.Error(), Category: ok}, nil
		}
		return GetCategoryResponse{Status: utils.StatusBadRequest, Message: "Category fetched Successfully", Category: ok}, err
	}
}

func makeListCategoriesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListCategories(ctx)
		return ListCategoriesResponse{Status: utils.StatusOK, Message: "Categories fetched successfully", Categories: ok}, err
	}
}
