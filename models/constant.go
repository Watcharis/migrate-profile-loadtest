package models

const (
	LIMIT_SIZE     = 1000
	TOTAL          = 10000
	GO_WORKER      = 10
	BUFFER_CHANNEL = 10
)

type ControlStep struct {
	Finally error
}
