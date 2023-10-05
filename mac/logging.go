package mac

import (
	"github.com/yuriykis/bth-speaker-on/log"
)

func Run(command string) (string, error) {
	log.Infof("Running command: %s", command)
	return run(command)
}

func Build(params ...string) string {
	log.Infof("Building command: %s", params)
	return build(params...)
}

func GetCurrentVolume() (v string) {
	defer log.Infof("Current volume is: %s", v)
	return getCurrentVolume()
}

func SetVolume(v float32) (vol string) {
	defer log.Infof("Set volume to: %s", vol)
	return setVolume(v)
}
