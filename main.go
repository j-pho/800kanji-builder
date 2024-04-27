package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode/utf8"
)

// $ go build
// $ ./400kanji-builder

func main() {
	splitter()
	// count_kanji()
}

func splitter() {
	const (
		F1  = "formatted-au-hke.txt"
		F2  = "formatted-e-khau.txt"
		F3  = "formatted-k-auhe.txt"
		F4  = "formatted-sent-au-khe.txt"
		F5  = "formatted-sent-k-auhe.txt"
		OUT = "merged-kanji-cards.txt"
	)

	files := []string{F1, F2, F3, F4, F5}
	ss := make([][]string, len(files))

	for i := range files {
		a, err := os.Open(files[i])
		if err != nil {
			panic(err)
		}
		defer a.Close()
		sc := bufio.NewScanner(a)
		for sc.Scan() {
			ss[i] = append(ss[i], sc.Text())
		}
	}

	f, err := os.Create(OUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buffer := bufio.NewWriter(f)
	for i := 0; i < len(ss[0]); i++ {
		for j := 0; j < len(ss); j++ {
			if i < len(ss[j]) {
				x := ss[j][i] + "\n"
				_, e := buffer.WriteString(x)
				if e != nil {
					log.Fatal(e)
				}
			}
		}
	}
}

//
// FYI
//

func count_kanji() {
	const (
		F1 = "./shell-scripts/kanji_for_audio.txt"
	)

	files := []string{F1}
	ss := make([][]string, len(files))

	for i := range files {
		a, err := os.Open(files[i])
		if err != nil {
			panic(err)
		}
		defer a.Close()
		sc := bufio.NewScanner(a)
		for sc.Scan() {
			ss[i] = append(ss[i], sc.Text())
		}
	}

	kset := make(map[rune]bool)
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(ss[i]); j++ {
			chh := ss[i][j]
			for len(chh) > 0 {
				r, size := utf8.DecodeRuneInString(chh)
				kset[r] = true
				chh = chh[size:]
			}
		}
	}
	klist := []rune{}
	for k := range kset {
		klist = append(klist, k)
	}

	sort.Sort(sortRunes(klist))
	fmt.Println(string(klist))
	fmt.Printf("%d distinct characters\n", len(kset))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
