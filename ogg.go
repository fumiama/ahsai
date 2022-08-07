package ahsai

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

func cutstream(s beep.StreamSeekCloser) {
	tmp := make([][2]float64, 1024)
	c := 0
	for c < 6 {
		_, _ = s.Stream(tmp)
		sum := (tmp[0][0] + tmp[0][1]) / 2
		for j := 1; j < 1024; j++ {
			sum += (tmp[j][0] + tmp[j][1]) / 2
			sum /= 2
		}
		if sum < 1e-32 && sum > -1e-32 {
			c++
		} else {
			c = 0
		}
	}
}

func SaveOggToFile(u, path string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	s, format, err := vorbis.Decode(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}
	defer s.Close()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	cutstream(s)
	return wav.Encode(f, s, format)
}

func SaveOggToWriteSeeker(u string, f io.WriteSeeker) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	s, format, err := vorbis.Decode(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}
	defer s.Close()
	cutstream(s)
	return wav.Encode(f, s, format)
}

func PlayOgg(u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	s, format, err := vorbis.Decode(resp.Body)
	if err != nil {
		resp.Body.Close()
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
