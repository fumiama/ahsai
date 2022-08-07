package ahsai

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const api = "https://cloud.ai-j.jp/demo/aitalk2webapi_nop.php?callback=callback&speaker_id=%d&text=%s&ext=ogg&volume=%.1f&speed=%.1f&pitch=%.1f&range=%.1f&anger=%.1f&sadness=%.1f&joy=%.1f&_=%d"

var (
	speakers = map[string]uint32{
		"琴葉葵":       551,
		"琴葉茜":       552,
		"紲星あかり":     554,
		"吉田くん":      1201,
		"東北ずん子":     1202,
		"月読アイ":      1203,
		"月読ショウタ":    1204,
		"民安ともえ":     1205,
		"結月ゆかり":     1206,
		"水奈瀬コウ":     1207,
		"京町セイカ":     1208,
		"東北きりたん":    1209,
		"桜乃そら":      1210,
		"東北イタコ":     1211,
		"ついなちゃん標準語": 1212,
		"ついなちゃん関西弁": 1213,
		"伊織弓鶴":      1214,
		"音街ウナ":      2006,
	}
)

type Speaker struct {
	id                                               uint32
	Volume, Speed, Pitch, Range, Anger, Sadness, Joy float32
}

func NewSpeaker() (s Speaker) {
	s.Volume = 1
	s.Speed = 1
	s.Pitch = 1
	s.Range = 1
	return
}

var (
	// ErrTextTooLong 文本超过 100 字
	ErrTextTooLong = errors.New("text too long")
	// ErrNoSuchSpeaker 查无此人
	ErrNoSuchSpeaker = errors.New("no such speaker")
)

func (s *Speaker) SetName(name string) error {
	id, ok := speakers[name]
	if !ok {
		return ErrNoSuchSpeaker
	}
	s.id = id
	return nil
}

// Speak text, return ogg url
func (s *Speaker) Speak(text string) (string, error) {
	if len([]rune(text)) > 100 {
		return "", ErrTextTooLong
	}
	resp, err := http.Get(fmt.Sprintf(api, s.id, url.QueryEscape(text), s.Volume, s.Speed, s.Pitch, s.Range, s.Anger, s.Sadness, s.Joy, time.Now().UnixMilli()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data = bytes.TrimPrefix(data, []byte(`callback({"url":"`))
	data = bytes.TrimSuffix(data, []byte(`"})`))
	data = bytes.ReplaceAll(data, []byte(`\/`), []byte(`/`))
	return "https:" + string(data), nil
}
