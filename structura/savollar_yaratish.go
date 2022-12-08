package main

import (
	"fmt"
)

type User struct {
	id                  uint8
	firstName, lastName string
}

type Option struct {
	id        uint8
	content   string
	isCorrect bool
}

type Question struct {
	id      uint8
	content string
	options []Option
}

// MVP - minimum viable product  samakat , vilosiped
func main() {
	var users = []User{
		{1, "Norbekov", "Shaxboz"},
		{2, "Shomurodov", "Jurabek"},
	}

	fmt.Println(users)
	var questions = []Question{
		{
			id:      1,
			content: "Qaysi dasturlash kursida tahsil olasiz?",
			options: []Option{
				{1, "Najot ta`lim", false},
				{2, "Uacademy", true},
				{3, "PDPacademy", false},
			},
		},
		{
			id:      2,
			content: "Qaysi dasturlash tilini bilasiz?",
			options: []Option{
				{1, "Python", false},
				{2, "Golang", true},
				{3, "C++", false},
			},
		},
		{
			id:      3,
			content: "Qaysi davlatda yashaysiz?",
			options: []Option{
				{1, "O`zbekiston", true},
				{2, "Tojikiston", false},
				{3, "Yaponiya", false},
			},
		},
		{
			id:      4,
			content: "Qaysi ta`lim musassasasida tahsil olgansiz?",
			options: []Option{
				{1, "Jahon tillari", false},
				{2, "MAdaniyat universiteti", false},
				{3, "Axborot texnologiyari universiteti", true},
			},
		},
		{
			id:      5,
			content: "Qaysi mobil kompaniyadan foydalanasiz?",
			options: []Option{
				{1, "samsung", true},
				{2, "apple", false},
				{3, "artel", false},
			},
		},
		{
			id:      6,
			content: "BLUE so`zining manosi?",
			options: []Option{
				{1, "Ko`k", true},
				{2, "Yashil", false},
				{3, "Sariq", false},
			},
		},
		{
			id:      7,
			content: "49sonining ildizini toping?",
			options: []Option{
				{1, "7", true},
				{2, "25", false},
				{3, "9", false},
			},
		},
		{
			id:      8,
			content: "Qaysi sport turiga qiziqasiz?",
			options: []Option{
				{1, "Boks", true},
				{2, "futbol", false},
				{3, "Karate", false},
			},
		},
		{
			id:      9,
			content: "Dunyoda nechta materik bor?",
			options: []Option{
				{1, "5", false},
				{2, "6", true},
				{3, "9", false},
			},
		},
		{
			id:      10,
			content: "Alisher Navoiy kim bo`lgan?",
			options: []Option{
				{1, "shoir", true},
				{2, "mulla", false},
				{3, "shoh", false},
			},
		},
	}

	for questions_index, question := range questions {
		fmt.Printf("%d. %s\n", questions_index+1, question.content)

		for options_index, option := range question.options {
			fmt.Printf("\t[%c]: %s\n", options_index+65, option.content)
		}
	}

	// fmt.Println(questions)

}
