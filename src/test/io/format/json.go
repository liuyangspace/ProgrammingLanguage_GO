package format

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func TestJsonEncoder()  {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}


func TestJsonDecoder()  {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if(err!=nil){
		println(err)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println("(string) "+k+" : "+ vv)
		case int:
			fmt.Println("(int) "+k+" : "+ string(vv))
		case float32,float64:
			fmt.Print("(float) "+k+" : ")
			fmt.Println(vv)
		case []interface{}:
			fmt.Println("(array) "+k+" : ")
			for i, u := range vv {
				fmt.Printf( "  %d : ",i)
				fmt.Println(u)
			}
		default:
			fmt.Println("(Unknow) "+k)
		}
	}
}