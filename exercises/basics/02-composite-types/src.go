package basics

import (
	"encoding/json"
	"sort"
	"strings"
)

func MustReturnAnArrayOfThreeIntegers() [3]int {
	a := [3]int{1, 2, 3}
	return a
}

func MustReturnTheFirstElementOfSlice(elements []int) int {
	// mySlice := make([]int, 3)
	// mySlice[0] = 1
	// mySlice[1] = 2
	// mySlice[2] = 3
	return elements[0]
}

func MustReturnTheLastElementOfSlice(elements []int) int {
	// mySlice := make([]int, 3)
	// mySlice[0] = 1
	// mySlice[1] = 2
	// mySlice[2] = 3
	return elements[len(elements)-1]
}

// When given ["My", "favorite", "language", "is", "Golang"], it should return [2, 8, 8, 2, 6]
func MustCountTheNumberOfLettersInEachWord(words []string) []int {

	nbOfLetters := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		nbOfLetters[i] = len(words[i])
	}

	return nbOfLetters
}

func MustReturnTheNumberOfWords(text string) int {

	if len(text) > 0 {
		//split string in words and assign it to a slice, then return the count the nb of element in slide
		s := strings.Split(text, " ")
		return len(s)
	} else {
		return 0
	}

}

func MustAppendTheWordGolang(words []string) []string {
	words = append(words, "Golang")
	return words
}

//if "level" key not exists, return 1, else return level.value
func MustReturnTheLevelKeyOrOneByDefault(stats map[string]int) int {

	levelValue, levelExist := stats["level"]

	if levelExist {
		return levelValue
	} else {
		return 1
	}

}

func MustSortAndReturn(listOfIntegers []int) []int {

	sort.Ints(listOfIntegers)

	return listOfIntegers
}

func MustReturn21YearsOldJohnDoe() Person {

	p := Person{"John", "Doe", 21}

	return p

}

// The function must panic if the json is invalid
// Learn more about the panic built-in here: https://gobyexample.com/panic
// You might want to read about Go and JSON: https://gobyexample.com/json
func MustReturnAListOfPerson(jsonListOfPerson []byte) []Person {

	//Marshal = serialize - Unmarshal = unserialize

	//dat = target = cible ds laquelle on renverse le json
	dat := make([]Person, 0)
	if err := json.Unmarshal(jsonListOfPerson, &dat); err != nil {
		panic(err)
	}

	return dat
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int
}
