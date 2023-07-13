package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title       string   `json:"title"`
	Genre       string   `json:"genre"`
	ReleaseYear int      `json:"releaseYear"`
	Director    Director `json:"director"`
}

type Director struct {
	Name        string `json:"name"`
	BirthYear   int    `json:"birthYear"`
	Nationality string `json:"nationality"`
}

type User struct {
	Name        string `json:"name,omitempty"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"active,omitempty"`
	LastLoginAt string `json:"lastLoginAt"`
}

func main() {
	var movie Movie
	var movies []User
	movieList := `[{"name":"John Doe","age":30,"email":"johndoe@example.com"}]`
	err := json.Unmarshal([]byte(movieList), &movies)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(movies)

	writer, err := os.Create("d:/phincon/weeks/week2/movie.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(movie)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Soal 5b
	reader, err2 := os.Open("d:/phincon/weeks/week2/movie.txt")
	if err2 != nil {
		fmt.Println(err.Error())
	}

	var movie2 Movie
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&movie2)
	if err != nil {
		fmt.Println(err.Error())
	}

	marshal, err := json.Marshal(movie2)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(marshal))

	user := `{"name":"John Doe","age":30,"email":"johndoe@example.com"}`
	var userTmp = struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}{}
	err = json.Unmarshal([]byte(user), &userTmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userTmp)

	user2 := User{Name: "Hakim", Age: 22, Active: true, LastLoginAt: "2023-05-06 22:00:00"}
	marshal, err = json.Marshal(user2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(marshal))

}
