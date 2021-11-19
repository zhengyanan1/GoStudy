package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string 			`json:"title"`
	Year int				`json:"year"`
	Price int				`json:"price"`
	Actors []string			`json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 1999, 30, []string{"xingye", "zhangbozhi"}}

//	编码过程 结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr;%s\n", jsonStr)

//	解码的过程，jsonstr ---> 结构体
// jsonstr:{"title":"喜剧之王","year":1999,"price3":30,"actors":["xingye","zhangbozhi"]}
	my_movie := Movie{}
//	my_movie := 9

	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil{
		fmt.Println("json unmarshal error", err)
	}

	fmt.Printf("decode: %v\n",my_movie)
}