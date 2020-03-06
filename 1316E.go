package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)
 
var dp [1 << 7]int64
var id []int
var a [100005]int
var s [100005][7]int
 
func main() {
	n := ReadInt()
	p := ReadInt()
	k := ReadInt()
	for i := 0; i < n; i++ {
		a[i] = ReadInt()
		id = append(id, i);
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			s[i][j] = ReadInt()
		}
	}
	sort.Slice(id,  func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})
	for i := 0; i < (1 << 7); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := (1 << uint(p)) - 1; j >= 0; j-- {
			var cnt int = 0
			for m := 0; m < p; m++ {
				if ((uint(j) >> uint(m)) & 1) != 0 {
					cnt++;
				}
			}
			cnt = i - cnt
			if cnt < k && dp[j] != -1 {
				dp[j] += int64(a[id[i]])
			}
			for m := 0; m < p; m++ {
				if ((uint(j) >> uint(m)) & 1) != 0 {
					if dp[j^(1<<uint(m))] != -1 {
						dp[j] = max(dp[j], dp[j^(1<<uint(m))]+int64(s[id[i]][m]))
					}
				}
			}
		}
	}
	fmt.Println(dp[(1<<uint(p))-1])
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