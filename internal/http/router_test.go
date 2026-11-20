package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	repomocks "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric/mocks"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	appErrors "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	httplib "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	serverHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/server"
	servermocks "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/server/mocks"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
)

func TestHTTPRouter(t *testing.T) {
	repo := repomocks.NewMockMetricsRepository(t)
	repo.EXPECT().ListMetrics().Return([]entity.Metric{
		{
			Name:         "requests",
			Type:         vo.CounterMetricType.String(),
			CounterValue: ptr(int64(42)),
		},
	}, nil)

	repo.EXPECT().FindMetricByNameAndType(mock.Anything, mock.Anything).
		Return(entity.Metric{}, appErrors.ErrNotFound)
	repo.EXPECT().CreateMetric(mock.Anything).Return(nil)
	repo.EXPECT().UpdateMetric(mock.Anything).Return(nil).Maybe()

	metricsUsecases := usecase.NewMetricUsecases(repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := servermocks.NewMockServerUsecase(ctrl)
	m.EXPECT().Ping().Return(nil).AnyTimes()

	metricsHandlers := metricHandlers.NewMetricHandler(metricsUsecases)
	serverHandlers := serverHandlers.NewServerHandler(m)

	handlers := httplib.HTTPHandlers{
		Metrics: metricsHandlers,
		Server:  serverHandlers,
	}

	router := httplib.NewHTTPRouter(handlers)
	ts := httptest.NewServer(router.Chi)
	defer ts.Close()

	tests := []struct {
		url         string
		method      string
		contentType string
		want        int
	}{
		{
			url:         "/",
			method:      http.MethodGet,
			contentType: "application/json",
			want:        http.StatusOK,
		},
		{
			url:         "/",
			method:      http.MethodPost,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodPut,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodDelete,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodPut,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodDelete,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/update/counter1/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/update/counter/some/123.123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusBadRequest,
		},
		{
			url:         "/update/gauge/some/123.123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/gauge/some/123.123sdf",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusBadRequest,
		},
		{
			url:         "/update/gauge/some/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/value/gauge/some",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/value/gauge/some",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			url := fmt.Sprintf("%s%s", ts.URL, tt.url)
			req, err := http.NewRequest(tt.method, url, nil)
			req.Header.Set("Content-Type", tt.contentType)
			require.NoError(t, err)
			w := httptest.NewRecorder()

			router.Chi.ServeHTTP(w, req)

			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

func TestHTTPRouter_Ping(t *testing.T) {
	tests := []struct {
		name       string
		pingErr    error
		wantStatus int
	}{
		{
			name:       "db ok",
			pingErr:    nil,
			wantStatus: http.StatusOK,
		},
		{
			name:       "db down",
			pingErr:    appErrors.ErrDatabaseConnection,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			serverUC := servermocks.NewMockServerUsecase(ctrl)
			serverUC.EXPECT().Ping().Return(tt.pingErr).Times(1)

			repo := repomocks.NewMockMetricsRepository(t)
			repo.EXPECT().ListMetrics().Return(nil, nil).Maybe()
			metricsUC := usecase.NewMetricUsecases(repo)

			router := httplib.NewHTTPRouter(httplib.HTTPHandlers{
				Metrics: metricHandlers.NewMetricHandler(metricsUC),
				Server:  serverHandlers.NewServerHandler(serverUC),
			})

			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			w := httptest.NewRecorder()
			router.Chi.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
