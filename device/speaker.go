package device

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/yuriykis/bth-speaker-on/log"

	"github.com/yuriykis/bth-speaker-on/mac"

	"github.com/andybrewer/mack"
)

type Speaker struct {
	Name string
}

func NewSpeaker(name string) *Speaker {
	return &Speaker{
		Name: name,
	}
}

func (s *Speaker) String() string {
	return fmt.Sprintf("Speaker")
}

func (s *Speaker) Up(wg *sync.WaitGroup) error {
	defer wg.Done()

	ok, err := s.musicPlaying()
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		if err := s.volume(0.2); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (s *Speaker) musicPlaying() (bool, error) {
	cmd1 := exec.Command("/usr/bin/pmset", "-g")
	cmd2 := exec.Command("grep", " sleep")
	output, err := cmd1.Output()
	if err != nil {
		return false, err
	}
	cmd2.Stdin = bytes.NewBuffer(output)
	output, err = cmd2.Output()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(output), "coreaudiod"), nil
}

func (s *Speaker) play() error {
	if err := mack.Say("Up"); err != nil {
		return err
	}
	log.Infof("Playing music")
	return nil
}

func (s *Speaker) volume(v float32) error {
	currentVol, err := mac.Run(mac.GetCurrentVolume())
	if err != nil {
		return err
	}

	_, err = mac.Run(mac.SetVolume(v))
	if err != nil {
		return err
	}
	if err := s.play(); err != nil {
		return err
	}

	currentVolFloat, err := strconv.ParseFloat(currentVol, 32)
	if err != nil {
		return err
	}
	_, err = mac.Run(mac.SetVolume(float32(currentVolFloat)))
	if err != nil {
		return err
	}
	return nil
}
