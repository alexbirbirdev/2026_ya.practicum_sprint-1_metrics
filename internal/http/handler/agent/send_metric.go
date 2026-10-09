package agent

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
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

func (h Handler) SendMetric(in dto.MetricDTO) (int, error) {
	url := fmt.Sprintf("%s/update/%s/%s/%v", h.baseURL, in.Type, in.Name, in.Value())
	resp, err := h.client.R().
		SetHeader("Content-Type", "text/plain").
		Post(url)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return resp.StatusCode(), nil
}
