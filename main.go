package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"log"
	"os"
)

const a = `[i][i:8][S:metadata][{][i][i:7][S:startAt][S][i][i:21][S:2022-08-28T15:51:13ZU][i][i:9][S:lastFrame][I][I:2011][i][i:7][S:players][{][i][i:1][S:0][{][i][i:10][S:characters][{][i][i:1][S:4][I][I:567][i][i:1][S:5][i][i:77][}][i][i:5][S:names][{][i][i:7][S:netplay][S][i][i:12][S:netplay-name][i][i:4][S:code][S][i][i:8][S:TEST#001][}][}][i][i:1][S:1][{][i][i:10][S:characters][{][i][i:1][S:1][i][i:123][}][}][}][}]`
const b = "7B69086D657461646174617B690773746172744174536915323032322D30382D32385431353A35313A31335A5569096C6173744672616D654907DB6907706C61796572737B6901307B690A636861726163746572737B690134490237690135694D7D69056E616D65737B69076E6574706C617953690C6E6574706C61792D6E616D656904636F646553690854455354233030317D7D6901317B690A636861726163746572737B690131697B7D7D7D7D7D"
const c = "7B69066669656C64315369036162637D"
const d = "[{][i][6][field1][S][i][3][abc][}]"

func main() {
	byt, err := ubjson.MarshalBlock(d)
	if err != nil {
		log.Fatal(err)
	}

	data, err := hex.DecodeString(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(byt)
	//marshalUBJSONTest()
	f, err := os.Open("replays/zelda-shiek.slp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bloo := bytes.NewReader(data)
	d := ubjson.NewDecoder(bloo)
	p := make(map[string]interface{})
	if err = d.Decode(&p); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
	//err = o.Decode(&players)
	//if err != nil {
	//	return errors.Wrap(err, "could not decode players into map")
	//}

	//decoder := ubjson.NewDecoder(f)
	//out := &slippi.Game{}
	//if err = decoder.Decode(out); err != nil {
	//	log.Fatal("could not decode into game: ", err)
	//}
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
