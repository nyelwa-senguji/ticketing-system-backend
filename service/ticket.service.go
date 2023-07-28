package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type CreateTicketRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int32  `json:"user_id"`
	CategoryID  int32  `json:"category_id"`
}

type ListTicketsResponse struct {
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Username     string `json:"username"`
	CategoryName string `json:"category_name"`
}

func (s service) CreateTicket(ctx context.Context, createTicketReq CreateTicketRequest) (string, error) {

	logger := log.With(s.logger, "method", "CreateTicket")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	ticket := db.CreateTicketParams{
		Subject:     createTicketReq.Subject,
		Description: createTicketReq.Description,
		Status:      createTicketReq.Status,
		UpdatedAt:   time,
		CreatedAt:   time,
		UserID:      createTicketReq.UserID,
		CategoryID:  createTicketReq.CategoryID,
	}

	_, err := s.repository.CreateTicket(ctx, ticket)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Ticket")

	return "Ticket created successfully", nil
}

func (s service) ListTickets(ctx context.Context) ([]ListTicketsResponse, error) {

	logger := log.With(s.logger, "method", "ListTickets")

	tickets, err := s.repository.ListTickets(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	var ListTickets []ListTicketsResponse

	for _, ticket := range tickets {

		GetUser, _ := s.repository.GetUser(ctx, ticket.UserID)

		GetCategory, _ := s.repository.GetCategory(ctx, ticket.CategoryID)

		t := ListTicketsResponse{
			Subject: ticket.Subject,
			Description: ticket.Description,
			Status: ticket.Status,
			Username: GetUser.Username,
			CategoryName: GetCategory.CategoryName,
		}

		ListTickets = append(ListTickets, t)

	}

	return ListTickets, nil
}
