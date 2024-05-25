package job

import (
	"fmt"
	http2 "net/http"
	"simple-job-board/internal/http"
	"simple-job-board/pkg/domain"
	"strconv"
	"time"
)

const (
	OfferIDRequireError   = "offer_id is required"
	OferrIDMustBeIntError = "offer_id must be an integer"
	OfferInvalidError     = "invalid offer %w"
)

type OfferHandler struct {
	service domain.JobService
}

func NewOfferHandler(service domain.JobService) *OfferHandler {
	return &OfferHandler{
		service: service,
	}
}

func (h *OfferHandler) GetOffer(w http2.ResponseWriter, r *http2.Request) {
	if r.PathValue("offer_id") == "" {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferIDRequireError))
		return
	}

	id, err := strconv.Atoi(r.PathValue("offer_id"))
	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OferrIDMustBeIntError))
		return
	}

	job, err := h.service.Get(r.Context(), id)
	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
		return
	}

	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
		return
	}
	http.ResponseJSON(w, http2.StatusOK, job)
}

func (h *OfferHandler) ListOffers(w http2.ResponseWriter, r *http2.Request) {
	params, err := http.SearchParams(r)
	if err != nil {
		http.ResponseJSON(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
		return
	}

	jobs, err := h.service.ListBy(r.Context(), params.Limit, params.Offset, true, "", "", 0, 0)
	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
		return
	}

	http.ResponseJSON(w, http2.StatusOK, jobs)
}

func (h *OfferHandler) CreateOffer(w http2.ResponseWriter, r *http2.Request) {
	formBody, err := CreateJobBody(r)
	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
	}

	_, err = h.service.Create(r.Context(), &domain.Job{
		Name:        formBody.Name,
		Description: formBody.Description,
		Level:       formBody.Level,
		SalaryMin:   formBody.SalaryMin,
		SalaryMax:   formBody.SalaryMax,
		EmployerID:  formBody.EmployerID,
		ValidDate:   time.Time{},
		Active:      true,
	})
	if err != nil {
		http.ResponseError(w, http2.StatusBadRequest, fmt.Errorf(OfferInvalidError, err))
		return
	}

	http.ResponseJSON(w, http2.StatusCreated, nil)
}
