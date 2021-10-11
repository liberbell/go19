package main

type Person struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Hobbies   string `json:"hobbies"`
}

func main() {
	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"Skiing", "Wind Surfing"}

}