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

> more info at [キャラクターで探す](https://www.ah-soft.com/product/chara.html)

- 琴葉葵
- 琴葉茜

![琴葉](/img/%E7%90%B4%E8%91%89.png)

- 紲星あかり

![紲星あかり](img/%E7%B4%B2%E6%98%9F%E3%81%82%E3%81%8B%E3%82%8A.png)

- 吉田くん

![吉田くん](img/%E5%90%89%E7%94%B0%E3%81%8F%E3%82%93.png)

- 東北ずん子

![東北ずん子](img/%E6%9D%B1%E5%8C%97%E3%81%9A%E3%82%93%E5%AD%90.png)

- 月読アイ

![月読アイ](img/%E6%9C%88%E8%AA%AD%E3%82%A2%E3%82%A4.png)

- 月読ショウタ

![月読ショウタ](img/%E6%9C%88%E8%AA%AD%E3%82%B7%E3%83%A7%E3%82%A6%E3%82%BF.png)

- 民安ともえ

![民安ともえ](img/%E6%B0%91%E5%AE%89%E3%81%A8%E3%82%82%E3%81%88.jpg)

- 結月ゆかり

![結月ゆかり](img/%E7%B5%90%E6%9C%88%E3%82%86%E3%81%8B%E3%82%8A.png)

- 水奈瀬コウ

![水奈瀬コウ](img/%E6%B0%B4%E5%A5%88%E7%80%AC%E3%82%B3%E3%82%A6.png)

- 京町セイカ

![京町セイカ](img/%E4%BA%AC%E7%94%BA%E3%82%BB%E3%82%A4%E3%82%AB.png)

- 東北きりたん

![東北きりたん](img/%E6%9D%B1%E5%8C%97%E3%81%8D%E3%82%8A%E3%81%9F%E3%82%93.png)

- 桜乃そら

![桜乃そら](img/%E6%A1%9C%E4%B9%83%E3%81%9D%E3%82%89.png)

- 東北イタコ

![東北イタコ](img/%E6%9D%B1%E5%8C%97%E3%82%A4%E3%82%BF%E3%82%B3.png)

- ついなちゃん標準語
- ついなちゃん関西弁

![ついなちゃん](img/%E3%81%A4%E3%81%84%E3%81%AA%E3%81%A1%E3%82%83%E3%82%93.png)

- 伊織弓鶴

![伊織弓鶴](img/%E4%BC%8A%E7%B9%94%E5%BC%93%E9%B6%B4.png)

- 音街ウナ

![音街ウナ](img/%E9%9F%B3%E8%A1%97%E3%82%A6%E3%83%8A.png)

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