package main

import (
	"fmt"
	"github.com/PMcca/go-slippi/slippi"
	"github.com/jmank88/ubjson"
	"io"
	"log"
	"os"
)

const asdasd = "7B690773746172744174536915323032322D30382D32385431353A35313A31335A5569096C6173744672616D65490BB86907706C61796572737B6901307B69056E616D65737B69076E6574706C617953690C4E6574706C6179204E616D656904636F646553690854455354233030317D690A636861726163746572737B690137490AF06902313955C87D7D6901317B69056E616D65737B69076E6574706C617953690E4E6574706C6179204E616D6520326904636F646553690854455354233030327D690A636861726163746572737B690133490BB87D7D7D7D"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ba() {
	gloo, err := os.Open("valid-meta.ubj")
	checkErr(err)

	defer gloo.Close()

	b, err := io.ReadAll(gloo)
	checkErr(err)
	m := slippi.Metadata{}
	if err = ubjson.Unmarshal(b, &m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
	//as, err := hex.DecodeString(asdasd)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := os.WriteFile("valid-meta.ubj", as, 0644); err != nil {
	//	log.Fatal(err)
	//}

	//reader := bytes.NewReader(as)
	//dec := ubjson.NewDecoder(reader)
	//p := make(map[string]interface{})
	//if err = dec.Decode(&p); err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(as))

	os.Exit(0)
}

func main() {
	ba()

	//byt, err := ubjson.MarshalBlock(d)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//data, err := hex.DecodeString(c)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(data)
	//fmt.Println(byt)
	//marshalUBJSONTest()
	f, err := os.Open("replays/zelda-shiek.slp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	asd, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(asd))

	d := ubjson.NewDecoder(f)
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
