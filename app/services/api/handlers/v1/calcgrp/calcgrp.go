package calcgrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Bruno-10/calc/business/core/calc"
	v1 "github.com/Bruno-10/calc/business/web/v1"
	"github.com/Bruno-10/calc/foundation/web"
)

// Handlers manages the set of calc endpoints.
type Handlers struct {
	calc *calc.Core
}

// New constructs a handlers for route access.
func New(calc *calc.Core) *Handlers {
	return &Handlers{
		calc: calc,
	}
}

func (h *Handlers) Execute(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var input string
	if err := web.Decode(r, &input); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	result, err := h.calc.Execute(ctx, input)
	if err != nil {
		return fmt.Errorf("execute: usr[%+v]: %w", result, err)
	}

	return web.Respond(ctx, w, result, http.StatusOK)

}
