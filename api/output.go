package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type MetricDefinition struct {
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	ImprovementDir int    `json:"improve_dir"`
}

type Metric struct {
	*MetricDefinition

	Value float64 `json:"value"`
}

type Event struct {
	RunEnv    *RunEnv `json:"context"`
	Timestamp int64   `json:"timestamp"`
	Metric    *Metric `json:"metric"`
}

func EmitMetric(ctx context.Context, def *MetricDefinition, value float64) {
	runenv := RunEnvFromContext(ctx)

	evt := &Event{
		RunEnv:    runenv,
		Timestamp: time.Now().UnixNano(),
		Metric:    &Metric{def, value},
	}

	bytes, err := json.Marshal(evt)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
