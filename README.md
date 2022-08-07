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