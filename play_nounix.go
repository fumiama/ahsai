//go:build android || darwin || js || windows
// +build android darwin js windows

package ahsai

import (
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// PlayOgg cut leading demo text and play directly
func PlayOgg(u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	s, format, err := vorbis.Decode(resp.Body)
	if err != nil {
		_ = resp.Body.Close()
		return err
	}
	defer s.Close()
	cutstream(s)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/32))
	if err != nil {
		return err
	}
	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- struct{}{}
	})))
	<-done
	return nil
}
