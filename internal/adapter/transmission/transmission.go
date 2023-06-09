package stock

import (
	"github.com/Goboolean/fetch-server/internal/domain/port"
	"github.com/Goboolean/fetch-server/internal/domain/value"
	"github.com/Goboolean/fetch-server/internal/infrastructure/prometheus"
	"github.com/Goboolean/shared/pkg/broker"
)

type TransmissionAdapter struct {
	conf broker.Configurator
	pub  broker.Publisher
}

func (a *TransmissionAdapter) TransmitStockBatch(tx port.Transactioner, stock string, batch []value.StockAggregate) error {

	prometheus.MQCounter.Add(float64(len(batch)))

	converted := make([]*broker.StockAggregate, len(batch))

	for idx := range converted {
		converted[idx] = &broker.StockAggregate{
			EventType: batch[idx].EventType,
			Average:   float32(batch[idx].Average),
			Min:       float32(batch[idx].Min),
			Max:       float32(batch[idx].Max),
			Start:     float32(batch[idx].Start),
			End:       float32(batch[idx].End),
			StartTime: batch[idx].StartTime,
			EndTime:   batch[idx].EndTime,
		}
	}

	return a.pub.SendDataBatch(stock, converted)
}

func (a *TransmissionAdapter) CreateStockQueue(tx port.Transactioner, stock string) error {
	return a.conf.CreateTopic(tx.Context(), stock)
}
