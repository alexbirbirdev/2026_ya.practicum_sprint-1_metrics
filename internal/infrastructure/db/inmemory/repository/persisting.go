package repository

import (
	"log"
	"time"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	interfaces "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	filestorage "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/file_storage"
)

var _ interfaces.MetricsRepository = (*PersistingMetricsRepository)(nil)

type PersistingMetricsRepository struct {
	innerRepo *MetricsRepository
	storage   filestorage.FileStorage
	interval  time.Duration
}

func NewPersistingMetricsRepository(
	innerRepo *MetricsRepository,
	storage filestorage.FileStorage,
	interval time.Duration,
) *PersistingMetricsRepository {
	return &PersistingMetricsRepository{
		innerRepo: innerRepo,
		storage:   storage,
		interval:  interval,
	}
}

func (p *PersistingMetricsRepository) saveAllMetrics() error {
	entities, err := p.innerRepo.ListMetrics()
	if err != nil {
		return err
	}
	dtos := make([]dto.Metrics, 0, len(entities))
	for _, entity := range entities {
		dtos = append(dtos, entity.MapToDTO())
	}
	return p.storage.Save(dtos)
}

func (p *PersistingMetricsRepository) syncSaveAllMetricsIfNeeded() {
	if p.interval != 0 {
		return
	}

	if err := p.saveAllMetrics(); err != nil {
		log.Printf("failed to save all metrics: %v", err)
	}
}

func (p *PersistingMetricsRepository) RunPeriodicMetricsSaving() {
	if p.interval <= 0 {
		return
	}

	go func() {
		for {
			if err := p.saveAllMetrics(); err != nil {
				log.Printf("failed to save all metrics periodically: %v", err)
			}

			time.Sleep(p.interval)
		}
	}()
}

func (p *PersistingMetricsRepository) CreateMetric(
	metric entity.Metric,
) error {
	if err := p.innerRepo.CreateMetric(metric); err != nil {
		return err
	}

	p.syncSaveAllMetricsIfNeeded()
	return nil
}

func (p *PersistingMetricsRepository) UpdateMetric(
	metric entity.Metric,
) error {
	if err := p.innerRepo.UpdateMetric(metric); err != nil {
		return err
	}

	p.syncSaveAllMetricsIfNeeded()
	return nil
}

func (p *PersistingMetricsRepository) ListMetrics() ([]entity.Metric, error) {
	return p.innerRepo.ListMetrics()
}

func (p *PersistingMetricsRepository) FindMetricByNameAndType(
	metricName string,
	metricType vo.MetricType,
) (entity.Metric, error) {
	return p.innerRepo.FindMetricByNameAndType(metricName, metricType)
}
