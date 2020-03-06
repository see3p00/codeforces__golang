package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
 
type Edge struct {
	to   int
	next int
}
 
var edge [200005 * 2] Edge
var head [200005] int
var edge1 [200005 * 2] Edge
var head1 [200005] int
var ans, cnt, cnt1, t, mi, ma int
var q [200005]int
var a [200005] int
var vis [200005] int
var dis [200005] int
 
func add(u, v int) {
	cnt++
	edge[cnt].to = v
	edge[cnt].next = head[u]
	head[u] = cnt
}
func add1(u, v int) {
	cnt1++
	edge1[cnt1].to = v
	edge1[cnt1].next = head1[u]
	head1[u] = cnt1
}
func bfs(s int, n int) {
	for i := 1; i <= n+1; i++ {
		dis[i] = 0x3f3f3f3f
		vis[i] = 0
	}
	t = 0
	dis[s] = 0
	vis[s] = 1
	q[t] = s;
	t++
	for i := 0; i < t; i++ {
		u := q[i]
		for j := head[u]; j != 0; j = edge[j].next {
			v := edge[j].to
			if dis[v] > dis[u]+1 {
				dis[v] = dis[u] + 1
				q[t] = v
				t++
			}
		}
	}
}
func main() {
	n := ReadInt()
	m := ReadInt()
	for i := 1; i <= m; i++ {
		u := ReadInt()
		v := ReadInt()
		add(v, u)
		add1(u, v)
	}
	k := ReadInt()
	for i := 1; i <= k; i++ {
		a[i] = ReadInt()
	}
	bfs(a[k], n)
	for i := 2; i <= k; i++ {
		if dis[a[i]] != dis[a[i-1]]-1 {
			mi++
			ma++
		} else {
			var cnt3 int = 0
			for j := head1[a[i-1]]; j != 0; j = edge1[j].next {
				if dis[edge1[j].to] == dis[a[i-1]]-1 {
					cnt3++
				}
			}
			if cnt3 >= 2 {
				ma++
			}
		}
	}
	fmt.Println(mi, ma)
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