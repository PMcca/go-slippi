package main

import (
	"fmt"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"log"
	"os"
)

func main() {
	//marshalUBJSONTest()
	f, err := os.Open("replays/4-players.slp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := ubjson.NewDecoder(f)
	out := &slippi.Game{}
	if err = decoder.Decode(out); err != nil {
		log.Fatal("could not decode into game: ", err)
	}
	x := 4
	fmt.Println(x)
	//
	//
	//res, err := io.ReadAll(f)
	//if err != nil {
	//	log.Fatal("could not read from file: ", err)
	//}
	////var out map[string]interface{}
	//out2 := &slippi.Game{}
	//
	//err = ubjson.Unmarshal(res, out2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//twoBytes := out2.Raw[3:5]
	//gameStartSize := binary.BigEndian.Uint16(twoBytes)
	//fmt.Println(gameStartSize)
	//
	////counts := map[byte]int{
	////	0x35: 0,
	////	0x36: 0,
	////	0x37: 0,
	////	0x38: 0,
	////	0x39: 0,
	////	0x3A: 0,
	////	0x3B: 0,
	////	0x3C: 0,
	////	0x3D: 0,
	////}
	////raw := out["raw"].([]byte)
	////for _, b := range raw {
	////	switch b {
	////	case
	////		0x35,
	////		0x36,
	////		0x37,
	////		0x38,
	////		0x39,
	////		0x3A,
	////		0x3B,
	////		0x3C,
	////		0x3D:
	////		counts[b] += 1
	////	default:
	////		continue
	////	}
	////}
	//
	////for k, v := range counts {
	////	fmt.Printf("\n%v: %d\n", k, v)
	////}
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
