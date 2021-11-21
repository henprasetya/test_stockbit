package main

import (
	"log"
	"sort"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

var Arraystr = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

func main() {

	loop(Arraystr)

}

func loop(array []string) {
	for _, s := range array {

		check(s, array)
		break
	}

}

func check(strcheck string, array []string) {
	var nextArray []string
	var macthAnagram []string
	for _, s := range array {

		eq := AreAnagram(s, strcheck)
		if !eq {
			nextArray = append(nextArray, s)
		}
		if eq {
			macthAnagram = append(macthAnagram, s)
		}
	}

	log.Print(macthAnagram)
	loop(nextArray)

}

func remove(slice []string, s int) []string {

	return append(slice[:s], slice[s+1:]...)
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func AreAnagram(s1, s2 string) bool {
	var r1 ByRune = StringToRuneSlice(s1)
	var r2 ByRune = StringToRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}
