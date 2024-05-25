package http

import (
	http2 "net/http"
	"simple-job-board/pkg/data"
	"strconv"
)

// Search Query Parameters

type searchParams struct {
	Limit  int `validate:"required,min=1"`
	Offset int `validate:"required,min=1"`
}

func SearchParams(r *http2.Request) (searchParams, error) {
	var p searchParams

	err := r.ParseForm()
	if err != nil {
		return p, err
	}

	limit, err := strconv.Atoi(r.Form.Get("limit"))
	if err != nil {
		return p, err
	}

	offset, err := strconv.Atoi(r.Form.Get("offset"))
	if err != nil {
		return p, err
	}

	p = searchParams{
		Limit:  limit,
		Offset: offset,
	}

	err = data.Valid(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}
