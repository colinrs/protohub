package infra

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/metric"
	"gorm.io/gorm"
)

const (
	gormClientNamespace = "gormc"
	metricsName         = "gormc_metrics_plugin"
)

const (
	CmdQuery  = "SELECT"
	CmdUpdate = "UPDATE"
	CmdInsert = "INSERT"
	CmdDelete = "DELETE"
	Unknown   = "unknown"
)

type ctxKey struct{}

var startTimeKey = ctxKey{}

type GormMetricsPlugin struct {
	queryCounter     metric.CounterVec
	errorCounter     metric.CounterVec
	latencyHistogram metric.HistogramVec
}

func NewGormMetricsPlugin() *GormMetricsPlugin {
	return &GormMetricsPlugin{
		queryCounter: metric.NewCounterVec(&metric.CounterVecOpts{
			Namespace: gormClientNamespace,
			Subsystem: "requests",
			Name:      "total",
			Help:      "GORM client requests count.",
			Labels:    []string{"db", "table", "operation"},
		}),
		errorCounter: metric.NewCounterVec(&metric.CounterVecOpts{
			Namespace: gormClientNamespace,
			Subsystem: "requests",
			Name:      "err_total",
			Help:      "GORM client requests err code count.",
			Labels:    []string{"db", "table", "operation"},
		}),
		latencyHistogram: metric.NewHistogramVec(&metric.HistogramVecOpts{
			Namespace: gormClientNamespace,
			Subsystem: "requests",
			Name:      "duration_ms",
			Help:      "GORM client requests duration(ms).",
			Labels:    []string{"db", "table", "operation"},
			Buckets:   []float64{0.25, 0.5, 1, 2, 5, 10, 25, 50, 100, 250, 500, 1000, 2000, 5000, 10000, 15000},
		}),
	}
}

func (p *GormMetricsPlugin) Name() string {
	return metricsName
}

func (p *GormMetricsPlugin) Initialize(db *gorm.DB) error {
	callbacks := []string{"create", "query", "update", "delete", "row"}
	for _, callback := range callbacks {
		err := db.Callback().Query().Before(fmt.Sprintf("gorm:%s", callback)).
			Register(fmt.Sprintf("%s:before_%s", metricsName, callback), p.before)
		if err != nil {
			return err
		}
		err = db.Callback().Query().After(fmt.Sprintf("gorm:%s", callback)).
			Register(fmt.Sprintf("%s:after_%s", metricsName, callback), p.after)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *GormMetricsPlugin) before(db *gorm.DB) {
	ctx := db.Statement.Context
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, startTimeKey, time.Now())
	db.Statement.Context = ctx
}

func (p *GormMetricsPlugin) after(db *gorm.DB) {
	ctx := db.Statement.Context
	startTime, ok := ctx.Value(startTimeKey).(time.Time)
	if !ok {
		// 如果没有找到开始时间，就使用当前时间（这种情况不应该发生）
		startTime = time.Now()
	}
	duration := time.Since(startTime)

	operation := matchOperationCommand(db.Statement.SQL.String())
	dbName := db.Name()
	tableName := db.Statement.Table

	p.queryCounter.Inc(dbName, tableName, operation)
	p.latencyHistogram.ObserveFloat(float64(duration.Milliseconds()), dbName, tableName, operation)

	if db.Error != nil {
		p.errorCounter.Inc(dbName, tableName, operation)
	}
}

func (p *GormMetricsPlugin) Close() error {
	return nil
}

func matchOperationCommand(sql string) string {
	if len(sql) > 6 {
		cmd := strings.ToUpper(sql[:6])
		switch cmd {
		case CmdQuery, CmdUpdate, CmdInsert, CmdDelete:
			return cmd
		}
	}
	return Unknown
}

func (p *GormMetricsPlugin) TranslateError(err error) error {
	return err
}
