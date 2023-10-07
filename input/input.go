package input

import (
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

type UpInterval struct {
	Value int
	Unit  time.Duration
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
		Unit:  time.Second, // TODO: make it configurable
	}

	if err != nil {
		return nil, err
	}
	return upInterval, nil
}
