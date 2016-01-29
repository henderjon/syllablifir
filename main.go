package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var numWords int

func init() {
	flag.IntVar(&numWords, "words", 4, "the number of words to output")
	flag.Parse()
}

func main() {
	ones := []string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t",
		"v", "w", "x", "z", "bl", "br", "ch", "cl", "cr", "dr", "fl", "fr", "gl", "gr",
		"pl", "pr", "sc", "sh", "sk", "sl", "sm", "sn", "sp", "st", "sw", "th", "tr",
		"tw", "wh", "wr", "sch", "scr", "shr", "sph", "spl", "spr", "squ", "str", "thr",
	}

	twos := []string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t",
		"v", "w", "x", "z", "ch", "sc", "sh", "sk", "st", "th", "sch", "sph",
	}

	vowels := []string{
		"a", "e", "i", "o", "u", "y", "ai", "au", "aw", "ay", "ea", "ee", "ei",
		"eu", "ew", "ey", "ie", "oi", "oo", "ou", "ow", "oy",
	}

	poss := len(ones) * len(vowels) * len(twos) * numWords

	log.Println(poss, "possibilities")

	letters := []*[]string{&ones, &twos, &vowels}
	for _, v := range letters {
		for i := 0; i < 5; i += 1 {
			*v = append(*v, *v...)
		}
	}

	for ; numWords > 0; numWords -= 1 {
		genWord(letters)
	}
}

func genWord(pieces []*[]string) {
	rand.Seed(time.Now().UnixNano())

	ones, twos, vowels := *pieces[0], *pieces[1], *pieces[2]
	st, mid, fin := rand.Intn(len(ones)), rand.Intn(len(vowels)), rand.Intn(len(twos))

	for st > len(ones) {
		st -= len(ones)
	}

	for mid > len(vowels) {
		mid -= len(vowels)
	}

	for fin > len(twos) {
		fin -= len(twos)
	}

	buf := bytes.Buffer{}

	buf.WriteString(ones[st])
	buf.WriteString(vowels[mid])
	buf.WriteString(twos[fin])
	buf.WriteString("\n")
	buf.WriteTo(os.Stdout)
}
