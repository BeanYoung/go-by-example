package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	b, _ := json.Marshal(true)
	fmt.Println(string(b))

	a, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(a))

	m, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(m))

	res1 := response1{1, []string{"a", "b", "c"}}
	r1, _ := json.Marshal(&res1)
	fmt.Println(string(r1))

	res2 := response2{1, []string{"a", "b"}}
	r2, _ := json.Marshal(&res2)
	fmt.Println(string(r2))

	byt := []byte(`{"page": 1, "fruits": ["a", "b", "c"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
	}
	fmt.Println(dat)
	fmt.Println(dat["page"].(float64))
	fruits := dat["fruits"].([]interface{})
	fmt.Println(fruits[0].(string))

	var res3 response2
	json.Unmarshal(byt, &res3)
	fmt.Println(res3)
}
