package agent

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	resty "github.com/go-resty/resty/v2"
)

type Handler struct {
	client  resty.Client
	baseURL string
}

func NewHandler(baseURL string) *Handler {
	return &Handler{
		client: *resty.New().
			SetRetryCount(3).
			SetRetryWaitTime(3 * time.Second).
			SetRetryMaxWaitTime(3 * time.Second),
		baseURL: baseURL,
	}
}

func (h Handler) SendMetric(in dto.Metrics) (int, error) {
	url := fmt.Sprintf("%s/update/", h.baseURL)

	resp, err := h.client.R().
		SetHeader("Content-Type", vo.ContentTypeApplicationJSON).
		SetBody(in).
		Post(url)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return resp.StatusCode(), nil
}
