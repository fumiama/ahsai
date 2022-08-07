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
