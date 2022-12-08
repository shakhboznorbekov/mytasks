package main

import "fmt"

func sonlar(a int) []string {
	var son []string
	var b, c int = 0, 0
	for a > 0 {
		c += 1
		b = a % 10
		a /= 10
		if c == 1 {
			birlik := map[int]string{
				0: "",
				1: "bir",
				2: "ikki",
				3: "uch",
				4: "to`rt",
				5: "besh",
				6: "olti",
				7: "yetti",
				8: "sakkiz",
				9: "to`qqiz",
			}
			son = append(son, birlik[b])
		} else if c == 2 {
			onlik := map[int]string{
				0: "",
				1: "o`n",
				2: "yigirma",
				3: "o`ttiz",
				4: "qirq",
				5: "ellik",
				6: "oltmish",
				7: "yetmish",
				8: "sakson",
				9: "to`qson",
			}
			son = append(son, onlik[b])
		} else if c == 3 {
			yuzlik := map[int]string{
				0: "",
				1: "bir yuz",
				2: "ikki yuz",
				3: "uch yuz",
				4: "to`rt yuz",
				5: "besh yuz",
				6: "olti yuz",
				7: "yetti yuz",
				8: "sakkiz yuz",
				9: "to`qqiz yuz",
			}
			son = append(son, yuzlik[b])
		} else if c == 4 {
			minglik := map[int]string{
				0: "",
				1: "bir ming",
				2: "ikki ming",
				3: "uch ming",
				4: "to`rt ming",
				5: "besh ming",
				6: "olti ming",
				7: "yetti ming",
				8: "sakkiz ming",
				9: "to`qqiz ming",
			}
			son = append(son, minglik[b])
		}

	}
	var out []string
	for i := len(son) - 1; i >= 0; i-- {
		out = append(out, son[i])
	}
	return out
}

func main() {
	var a int = 9101
	fmt.Println(sonlar(a))
}
