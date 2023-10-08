package input

import (
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

const (
	DefaultUpIntervalValue = 5
)

type UpInterval struct {
	Value int
	Unit  time.Duration
}

func DefaultUpInterval() *UpInterval {
	return &UpInterval{
		Value: DefaultUpIntervalValue,
		Unit:  time.Minute,
	}
}

func (i *UpInterval) Duration() time.Duration {
	return time.Duration(i.Value) * i.Unit
}

func ParseUpIntervalFlag(fv *pflag.Flag) (*UpInterval, error) {
	fvString := fv.Value.String()
	fvInt, err := strconv.Atoi(fvString)
	if err != nil {
		return nil, err
	}
	upInterval := &UpInterval{
		Value: fvInt,
		Unit:  time.Minute, // TODO: make it configurable
	}

	if err != nil {
		return nil, err
	}
	return upInterval, nil
}
