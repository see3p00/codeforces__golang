package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
var a [100005]int
func main() {
	//var n, d, t, i, cas int
	for cas := ReadInt(); cas >0; cas-- {
		var maxx, ans int
		n:=ReadInt()
		d:=ReadInt()
		for i := 1; i <= n; i++ {
			a[i]=ReadInt()
			maxx = max(a[i], maxx)
			if a[i] == d {
				ans = 1
			}
		}
		if ans == 0 {
			tmp  := (d + maxx - 1) / maxx
			ans = max(tmp, 2)
		}
		fmt.Println(ans)
	}
}
func max(a, b int) int {
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