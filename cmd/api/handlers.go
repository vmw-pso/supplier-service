package main

import (
	"net/http"

	"github.com/vmw-pso/toolkit"
)

func (s *server) handleSupplierNames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		suppliers, err := s.Models.Supplier.GetAll(s.DB)
		if err != nil {
			s.tools.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		payload := toolkit.JSONResponse{
			Error:   false,
			Message: "All suppliers",
			Data:    suppliers,
		}

		_ = s.tools.WriteJSON(w, http.StatusAccepted, payload)
	}
}
