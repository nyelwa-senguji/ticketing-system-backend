package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
)

func decodeAssignPermissionToRoleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.AssignPermissionToRoleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeListAssignedPermissionToRoleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 32)
	id := int32(i)
	request := endpoint.ListAssignedPermissionsToRoleRequest{
		RoleID: id,
	}
	return request, nil
}

func decodeRevokePermissionToRoleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req db.RevokePermissionRoleParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

