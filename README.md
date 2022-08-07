# ahsai
AH Soft フリーテキスト音声合成 demo API

## demo
Just run go test to hear the voice below

<audio src='/test.wav' controls><a href='/test.wav'>こんにちは、世界</a></audio>

```go
package ahsai

import "testing"

func TestAPI(t *testing.T) {
	s := NewSpeaker()
	err := s.SetName("東北イタコ")
	if err != nil {
		t.Fatal(err)
	}
	u, err := s.Speak("こんにちは、世界")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
	err = PlayOgg(u)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveOggToFile(u, "test.wav")
	if err != nil {
		t.Fatal(err)
	}
}
```

## supported speakers
- 琴葉葵
- 琴葉茜
- 紲星あかり
- 吉田くん
- 東北ずん子
- 月読アイ
- 月読ショウタ
- 民安ともえ
- 結月ゆかり
- 水奈瀬コウ
- 京町セイカ
- 東北きりたん
- 桜乃そら
- 東北イタコ
- ついなちゃん標準語
- ついなちゃん関西弁
- 伊織弓鶴
- 音街ウナ

## commandline tool
```bash
go run cmd/main.go -h
Usage:
  -a float
        anger
  -b uint
        border slience sample lenth (default 2048)
  -d float
        sadness
  -f string
        line-separated text to read
  -h    display this help
  -j float
        joy
  -n string
        specify speaker (default "民安ともえ")
  -o string
        output wav file path (default "out.wav")
  -p float
        pitch (default 1)
  -r float
        range (default 1)
  -s float
        speed (default 1)
  -v float
        volume (default 1)
```