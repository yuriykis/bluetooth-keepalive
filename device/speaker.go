package device

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

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
	log.Println("Playing music")
	return nil
}

func (s *Speaker) volume(v float32) error {
	cmd1 := exec.Command("osascript", "-e", "output volume of (get volume settings)")
	output, err := cmd1.Output()
	if err != nil {
		return err
	}
	currentVol := string(output)

	cmd2 := exec.Command("osascript", "-e", fmt.Sprintf("set volume output volume %f", v))
	log.Println("Setting volume to", v)
	if err = cmd2.Run(); err != nil {
		return err
	}

	if err := s.play(); err != nil {
		return err
	}

	cmd3 := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf("set volume output volume %s", currentVol),
	)
	log.Println("Setting volume back to", currentVol)
	if err = cmd3.Run(); err != nil {
		return err
	}
	return nil
}
