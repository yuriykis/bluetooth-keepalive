package system

import (
	"context"
	"time"

	"github.com/yuriykis/bth-speaker-on/device"
	"github.com/yuriykis/bth-speaker-on/log"
)

type LoggingDeviceManagerMiddleware struct {
	next DeviceManager
}

func NewLoggingDeviceManagerMiddleware(
	next DeviceManager,
) DeviceManager {
	return &LoggingDeviceManagerMiddleware{
		next: next,
	}
}

func (ldmm *LoggingDeviceManagerMiddleware) Devices() (devices []device.Devicer, err error) {
	log.WithFields(log.Fields{
		"devices": devices,
		"err":     err,
	}).Info("LoggingDeviceManagerMiddleware.Devices()")
	return ldmm.next.Devices()
}

func (ldmm *LoggingDeviceManagerMiddleware) Start(
	ctx context.Context,
	d time.Duration,
) error {

	log.WithFields(log.Fields{
		"duration": d,
	}).Info("Starting DeviceManager...")

	defer func() {
		log.Info("Exiting DeviceManager...")
	}()

	return ldmm.next.Start(ctx, d)
}
