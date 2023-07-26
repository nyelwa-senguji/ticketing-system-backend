package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
	Status       string `json:"status"`
}

func (s service) CreateCategory(ctx context.Context, createCategoryReq CreateCategoryRequest) (string, error) {

	logger := log.With(s.logger, "method", "CreateCategory")

	checkIfCategoryExists, _ := s.repository.GetCategoryByName(ctx, createCategoryReq.CategoryName)

	if checkIfCategoryExists.CategoryName == createCategoryReq.CategoryName {
		return "Category already exists", nil
	}

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	category := db.CreateCategoryParams{
		CategoryName: createCategoryReq.CategoryName,
		Status:       createCategoryReq.Status,
		CreatedAt:    time,
		UpdatedAt:    time,
	}

	_, err := s.repository.CreateCategory(ctx, category)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Role")

	return "Category created successfully", nil
}

func (s service) GetCategory(ctx context.Context, id int32)(db.Category, error){
	logger := log.With(s.logger, "method", "GetCategory")

	category, err := s.repository.GetCategory(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return category, err
	}

	logger.Log("Get Category", category.CategoryName)

	return category, nil
}

func (s service) ListCategories(ctx context.Context) ([]db.Category, error){
	logger := log.With(s.logger, "method", "ListCategories")

	categories, err := s.repository.ListCategories(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("List All Categories")

	return categories, nil
}
