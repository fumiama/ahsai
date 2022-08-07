package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/faiface/beep/wav"
	"github.com/fumiama/ahsai"
)

func main() {
	f := flag.String("f", "", "line-separated text to read")
	o := flag.String("o", "out.wav", "output wav file path")
	n := flag.String("n", "民安ともえ", "specify speaker")
	v := flag.Float64("v", 1.0, "volume")
	s := flag.Float64("s", 1.0, "speed")
	p := flag.Float64("p", 1.0, "pitch")
	r := flag.Float64("r", 1.0, "range")
	a := flag.Float64("a", 0, "anger")
	d := flag.Float64("d", 0, "sadness")
	j := flag.Float64("j", 0, "joy")
	b := flag.Uint("b", 2048, "border slience sample lenth")
	h := flag.Bool("h", false, "display this help")
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	if *f == "" {
		panic("parameter -f must be specified")
	}
	spk := ahsai.Speaker{Volume: float32(*v), Speed: float32(*s), Pitch: float32(*p), Range: float32(*r), Anger: float32(*a), Sadness: float32(*d), Joy: float32(*j)}
	err := spk.SetName(*n)
	if err != nil {
		panic(err)
	}
	txt, err := os.Open(*f)
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(txt)
	lst := make([]string, 0, 128)
	i := 0
	for sc.Scan() {
		t := sc.Text()
		i++
		if len([]rune(t)) > 100 {
			panic("line " + strconv.Itoa(i) + ": too long (> 100 chars)")
		}
		lst = append(lst, t)
	}
	err = txt.Close()
	if err != nil {
		panic(err)
	}
	for i, t := range lst {
		u, err := spk.Speak(t)
		if err != nil {
			panic("line " + strconv.Itoa(i) + "error: " + err.Error())
		}
		lst[i] = u
		fmt.Print("\rread: ", (i+1)*100/len(lst), " %")
	}
	sm, format, err := ahsai.ComposeStream(*b, func(p int) { fmt.Print("\rcompose: ", p, " %") }, lst...)
	if err != nil {
		panic(err)
	}
	out, err := os.Create(*o)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	err = wav.Encode(out, sm, format)
	if err != nil {
		panic(err)
	}
	fmt.Println("\rall process succeed")
}
