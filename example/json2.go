package main

import (
    "fmt"
    "encoding/json"
)

type interest struct {
    Name string `json:"sport"`
    AttendNum int
}

type user struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Interesting *interest
    gender string
}



func main() {
    obj := &user{
	Id: 1,
	Name: "frankie",
	Age: 18,
        gender: "M"}
    jsonObj, _ := json.Marshal(obj)
    fmt.Println(obj)
    fmt.Println(jsonObj)
    fmt.Println(string(jsonObj))

    inObj := new(interest)
    inObj.Name = "swimming"
    inObj.AttendNum = 1
    obj.Interesting = inObj
    jsonObj, _ = json.Marshal(obj)
    fmt.Println(string(jsonObj))

    str := `{"id": 2, "gender": "F", "name": "frank", "age": 22, "Interesting": {"sport":"jogging", "AttendNum": 1}}`
    byt := []byte(str)
    fmt.Println(str)
    fmt.Println(byt)
    //var obj2 user
    obj2 := user{}
    if err := json.Unmarshal(byt, &obj2); err != nil {
        panic("oh, no")
    }
    fmt.Println(obj2) 
    fmt.Println(obj2.Age) 
    fmt.Println(obj2.Interesting) 
    fmt.Println(obj2.Interesting.AttendNum) 
}

