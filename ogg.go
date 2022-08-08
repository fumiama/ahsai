package ahsai

import (
	"io"
	"net/http"
	"os"

	"github.com/faiface/beep"
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

// ComposeStream 组合 urls 成为单个 stream, 并留出 sil 采样间隔
func ComposeStream(sil uint, progress func(p int), urls ...string) (sm beep.Streamer, format beep.Format, err error) {
	var buf *beep.Buffer
	for i, u := range urls {
		var resp *http.Response
		resp, err = http.Get(u)
		if err != nil {
			return
		}
		var s beep.StreamSeekCloser
		s, format, err = vorbis.Decode(resp.Body)
		if err != nil {
			_ = resp.Body.Close()
			return
		}
		if i == 0 {
			buf = beep.NewBuffer(format)
		}
		cutstream(s)
		buf.Append(beep.Silence(int(sil)))
		buf.Append(s)
		_ = s.Close()
		if progress != nil {
			progress((i + 1) * 100 / len(urls))
		}
	}
	sm = buf.Streamer(0, buf.Len())
	return
}

// SaveOggToFile cut leading demo text and save wav to path
func SaveOggToFile(u, path string) error {
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
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	cutstream(s)
	return wav.Encode(f, s, format)
}

// SaveOggToWriteSeeker cut leading demo text and write wav stream to f
func SaveOggToWriteSeeker(u string, f io.WriteSeeker) error {
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
	return wav.Encode(f, s, format)
}
