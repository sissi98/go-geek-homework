package web

import (
	"fmt"
	"regexp"
	"testing"
)

func TestComplie(t *testing.T) {
	tmp, e := regexp.Compile(".*")
	fmt.Println(tmp)
	fmt.Println(e)
}

func TestMatch(t *testing.T) {
	tmp, _ := regexp.Compile("^[0-9]+$")
	ok := tmp.Match([]byte("12434"))
	ok1 := tmp.MatchString("12")
	fmt.Println(ok)
	fmt.Println(ok1)
}


func TestMatchq(t *testing.T) {
	tmp, _ := regexp.Compile("[0-9]+")
	ok := tmp.Match([]byte("abc"))
	
	fmt.Println(ok)
	
}
