package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"
)

const defaultUrlSize = 8

type register struct {
	timestamp  string
	reducedUrl string
}

func main() {
	db := make(map[string]register)
	for {
		command := Menu()
		switch command {
		case 1:
			fmt.Print("Type URL: ")
			var input string
			// input = "https://www.google.com"
			fmt.Scan(&input)
			result := Execute(input, db)
			fmt.Println("")
			fmt.Println("Reduced URL:", result.reducedUrl, " created on:", result.timestamp)
		case 2:
			PrintAllUrls(db)
		case 0:
			fmt.Println("Bye!!")
			os.Exit(-1)
		}
	}
}

func Menu() int {
	fmt.Println("")
	fmt.Println("1 - Reduce URL")
	fmt.Println("2 - Check all reduced URLs")
	fmt.Println("0 - Exit")
	fmt.Println("")

	var command int
	fmt.Scan(&command)
	// command = 1
	return command
}

func Execute(input string, db map[string]register) register {
	register, found := cache(input, db)
	if found {
		return register
	} else {
		newRegister := ReduceUrl(input, db)
		db[input] = newRegister
		return newRegister
	}
}

func cache(input string, db map[string]register) (register, bool) {
	_, found := db[input]
	if found {
		return db[input], true
	} else {
		return register{}, false
	}
}

func ReduceUrl(input string, db map[string]register) register {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	var newUrl string
	for {
		newUrl = GenerateNewUrl(u)
		if IsValidUrl(newUrl, db) {
			break
		}
	}

	t := time.Now().Format(time.RFC822)
	return register{t, newUrl}
}

func GenerateNewUrl(rawurl *url.URL) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, defaultUrlSize)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	newUrl := rawurl.Scheme + "://" + string(s)
	return newUrl
}

func IsValidUrl(newUrl string, db map[string]register) bool {
	result := true
	for _, v := range db {
		if v.reducedUrl == newUrl {
			result = false
		}
	}
	return result
}

func PrintAllUrls(db map[string]register) {
	for k, v := range db {
		fmt.Println(k, " -> ", v.reducedUrl, " Generated on:", v.timestamp)
	}
}
