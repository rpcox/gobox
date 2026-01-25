package main

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type JsonStruct struct {
	Key1 string  `json:"key1"`
	Key2 int64   `json:"key2"`
	Key3 float64 `json:"key3"`
	Key4 JsonSub1 `json:"key4"`
}

type JsonSub1 struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

var jsi = jsoniter.ConfigCompatibleWithStandardLibrary

func MarshalStd(js *JsonStruct) ([]byte, error) {
	data, err := json.Marshal(*js)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func MarshalJsi(js *JsonStruct) ([]byte, error) {
	data, err := jsi.Marshal(*js)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UnMarshalStd(data []byte) (JsonStruct, error) {
	var js JsonStruct = JsonStruct{}
	err := json.Unmarshal(data, &js)
	if err != nil {
		return JsonStruct{}, err
	}

	return js, nil
}

func UnMarshalJsi(data []byte) (JsonStruct, error) {
	var js JsonStruct = JsonStruct{}
	err := jsi.Unmarshal(data, &js)
	if err != nil {
		return JsonStruct{}, err
	}

	return js, nil
}

func main() {
	js := JsonStruct{
		Key1: "Value1", 
		Key2: 1 << 32, 
		Key3: 3.14, 
		Key4: JsonSub1{
			Length: 5, 
			Width: 6,
		},
	}

	data, err := MarshalStd(&js)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	js2, err := UnMarshalStd(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(js2)

	data, err = MarshalJsi(&js)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	js3, err := UnMarshalJsi(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(js3)

}
	
