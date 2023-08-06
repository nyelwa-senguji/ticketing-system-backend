package endpoint

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type (
	CreateTicketResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ListTicketsResponse struct {
		Status  int                           `json:"status"`
		Message string                        `json:"message"`
		Tickets []service.ListTicketsResponse `json:"tickets"`
	}
)

func makeCreateTicketEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreateTicketRequest)

		if reflect.DeepEqual(req.Subject, "") {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: "Subject cannot be empty", Success: false}, nil
		}

		if reflect.DeepEqual(req.UserID, "") {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: "User cannot be empty", Success: false}, nil
		}

		if reflect.DeepEqual(req.CategoryID, "") {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: "Category cannot be empty", Success: false}, nil
		}

		ok, err := s.CreateTicket(ctx, req)
		if err != nil {
			return CreateCategoryResponse{Status: utils.StatusBadRequest, Message: err.Error(), Success: false}, nil
		}

		return CreateCategoryResponse{Status: utils.StatusOK, Message: ok, Success: true}, err
	}
}

func makeListTicketsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListTickets(ctx)
		return ListTicketsResponse{Status: utils.StatusOK, Message: "Tickets fetched successfully", Tickets: ok}, err
	}
}
