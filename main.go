package main

import (
	"bufio"
	"fmt"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"io"
	"log"
	"os"
)

func main() {
	//marshalUBJSONTest()
	f, err := os.Open("replays/test.slp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)
	var res []byte

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
		res = append(res, buf...)
	}

	var out map[string]interface{}
	out2 := &slippi.Game{}

	err = ubjson.Unmarshal(res, out2)
	if err != nil {
		log.Fatal(err)
	}

	counts := map[byte]int{
		0x35: 0,
		0x36: 0,
		0x37: 0,
		0x38: 0,
		0x39: 0,
		0x3A: 0,
		0x3B: 0,
		0x3C: 0,
		0x3D: 0,
	}
	raw := out["raw"].([]byte)
	for _, b := range raw {
		switch b {
		case
			0x35,
			0x36,
			0x37,
			0x38,
			0x39,
			0x3A,
			0x3B,
			0x3C,
			0x3D:
			counts[b] += 1
		default:
			continue
		}
	}

	for k, v := range counts {
		fmt.Printf("\n%v: %d\n", k, v)
	}
}

func marshalUBJSONTest() {
	g := slippi.Game2{
		FieldA: "hello",
		FieldB: "world",
	}

	b, err := ubjson.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	fmt.Println(s)

	out := slippi.Game3{}
	err = ubjson.Unmarshal(b, &out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
	os.Exit(0)
}
