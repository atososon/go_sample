package main

import (
	"fmt"
	"bytes"
	"errors"
	"flag"
	"os"
	"bufio"
	"strconv"
	"strings"
)
type Board [9][9]int

func pretty(b Board) string{
	var buf bytes.Buffer
	for i:=0; i<9; i++{
		if i%3 == 0{
			buf.WriteString("+---+---+---+\n")
		}
		for j:=0; j<9; j++{
			if j%3 == 0{
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func duplicated(c [10]int) bool{
	for k, v := range c{
		if k == 0{
			continue
		}
		if v >= 2{
			return true
		}
	}
	return false
}

func verify(b Board) bool{
	for i:=0; i<9; i++{
		var c [10]int
		for j:=0; j<9; j++{
			c[b[i][j]]++
		}
		if duplicated(c){
			return false
		}
	}
	for i:=0; i<9; i++{
		var c [10]int
		for j:=0; j<9; j++{
			c[b[j][i]]++
		}
		if duplicated(c){
			return false
		}
	}
	for i:=0; i<9; i+=3{
		for j:=0; j<9; j+=3{
			var c [10]int
			for row:=i; row<i+3; row++{
				for col:=j; col<j+3; col++{
					c[b[row][col]]++
				}
			}
			if duplicated(c){
				return false
			}
		}
	}
	return true
}

func solved(b Board) bool{
	if !verify(b){
		return false
	}
	for i:=0; i<9; i++{
		for j:=0; j<9; j++{
			if b[i][j] == 0{
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board) bool{
	if solved(*b){
		return true
	}
	for i:=0; i<9; i++{
		for j:=0; j<9; j++{
			if b[i][j] == 0{
				for c:=9; c>=1; c--{
					b[i][j] = c
					if verify(*b){
						if backtrack(b){
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func short(input string) (*Board, error){
	if len(input) != 81{
		return nil, errors.New("input short stringlength must be 81")
	}
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanRunes)
	var b Board
	for i:=0; i<9; i++{
		for j:=0; j<9; j++{
			if !s.Scan(){
				break
			}
			token := s.Text()
			if token == "."{
				b[i][j] = 0
				continue
			}
			n, err := strconv.Atoi(token)
			if err != nil{
				return nil, err
			}
			b[i][j] = n
		}
	}
	return &b, nil
}

func main(){
	flag.Parse()
	input := flag.Arg(0)
	fmt.Printf("input = %+v\n", input)
	b, err := short(input)
	if err != nil{
		panic(err)
	}
	if backtrack(b){
		fmt.Println(pretty(*b))
	}else{
		fmt.Fprintf(os.Stderr, "cannot solve")
	}
}