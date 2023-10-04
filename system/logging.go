package system

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yuriykis/bth-speaker-on/device"
)

var logger *logrus.Logger

type LoggingDeviceManagerMiddleware struct {
	next DeviceManager
}

func NewLoggingDeviceManagerMiddleware(
	next DeviceManager,
	log *logrus.Logger,
) DeviceManager {
	logger = log
	return &LoggingDeviceManagerMiddleware{
		next: next,
	}
}

func (ldmm *LoggingDeviceManagerMiddleware) Devices() (devices []device.Devicer, err error) {
	logger.WithFields(logrus.Fields{
		"devices": devices,
		"err":     err,
	}).Info("LoggingDeviceManagerMiddleware.Devices()")
	return ldmm.next.Devices()
}

func (ldmm *LoggingDeviceManagerMiddleware) Start(
	ctx context.Context,
	d time.Duration,
) error {

	logger.WithFields(logrus.Fields{
		"duration": d,
	}).Info("Starting DeviceManager...")

	defer func() {
		logger.Info("Exiting DeviceManager...")
	}()

	return ldmm.next.Start(ctx, d)
}
