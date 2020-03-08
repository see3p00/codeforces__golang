package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var a [40005]int
var b [40005]int
var c [100005]int

func main() {
	n := ReadInt()
	m := ReadInt()
	k := ReadInt()
	cnt := 0
	var ans int64
	for i := 1; i <= n; i++ {
		a[i] = ReadInt()
	}
	for i := 1; i <= m; i++ {
		b[i] = ReadInt()
	}
	for i := 1; i*i <= k; i++ {
		if k%i == 0 {
			c[cnt] = i
			cnt++
			if i == k/i {
				continue
			}
			c[cnt] = k / i
			cnt++
		}
	}
	for i := 0; i < cnt; i++ {
		tmp := 0
		var x, y int64
		for j := 1; j <= n; j++ {
			if a[j] == 1 {
				tmp++
			} else {
				tmp = 0
			}
			if tmp == c[i] {
				tmp--
				x++
			}
		}
		tmp = 0
		for j := 1; j <= m; j++ {
			if b[j] == 1 {
				tmp++
			} else {
				tmp = 0
			}
			if tmp == k/c[i] {
				tmp--
				y++
			}
		}
		ans += x * y
	}
	fmt.Println(ans)
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
