package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
 
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
 
func main() {
	for t := ReadInt(); t > 0; t-- {
		n := ReadInt()
		s := ReadString()
		ans := s
		var k int =0
		for i := 0; i < n; i++ {
			tmp1 := s[i:n]
			tmp2 := s[0:i]
			if (n-i)%2 == 1 {
				tmp2=reverseString(tmp2)
			}
			tmp1 += tmp2
			if tmp1 < ans {
				ans = tmp1
				k = i
			}
		}
		fmt.Println(ans)
		fmt.Println(k + 1)
	}
}
 
////////////////////////////////////
 
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
 
var out chan string
var in *bufio.Scanner
var outWg *sync.WaitGroup
 
func ReadString() string {
	in.Scan()
	return in.Text()
}
func ReadStringSlice(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = ReadString()
	}
	return s
}
func ReadInt() int {
	intStr := ReadString()
	i, _ := strconv.ParseInt(intStr, 10, 32)
	return int(i)
}
func ReadIntSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = ReadInt()
	}
	return arr
}
func init() {
	//set input
	in = bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1024), int(2e+5))
	in.Split(bufio.ScanWords)
 
	//set output
	out = make(chan string, 16)
	outWg = &sync.WaitGroup{}
	outWg.Add(1)
 
	writer := bufio.NewWriterSize(os.Stdout, int(2e+5))
	go func(write *bufio.Writer) {
		defer outWg.Done()
		defer write.Flush()
 
		for line := range out {
			write.WriteString(line + "\n")
		}
	}(writer)
}