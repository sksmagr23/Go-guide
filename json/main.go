package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"` // we can define aliases for json using ``
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`      /* it tells to not show this field in json */
	Tags []string `json:"tags,omitempty"` // omitempty here, will do that if field is nil, it will not show the field
}

func main()  {
	encodeJson()
	DecodeJson()
}

func encodeJson()  {
	myCourses := []course{
		{"ReactJS", 300, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"MERN", 350, "lco.in", "abc1234", []string{"web-dev", "mongo"}},
		{"GOlang", 200, "lco.in", "xyz123",nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t") // \t indicates tab, second arg is prefix , which can be usually empty
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson()  {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "ReactJS",
        "Price": 300,
        "website": "lco.go",
        "tags": ["web-dev","js"]
    }
	`)

	var mycour course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &mycour)
		fmt.Printf("%#v\n", mycour)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where we just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}