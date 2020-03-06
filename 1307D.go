package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)
 
type Edge struct {
	to   int
	next int
}
 
var edge [200005 * 2] Edge
//var a [200005] int
var head [200005] int
var ans, cnt, t int
var q [200005]int
var vis [200005] int
var dis1 [200005] int
var dis2 [200005] int
 
func add(u, v int) {
	cnt++
	edge[cnt].to = v
	edge[cnt].next = head[u]
	head[u] = cnt
}
func bfs(dis *[200005]int,s int, n int) {
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
type shuzu []int
 
func (a shuzu) Len() int {
	return len(a)
}
func (a shuzu) Less(i, j int) bool {
	return dis1[a[i]] < dis1[a[j]]
}
func (a shuzu) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func main() {
	n := ReadInt()
	m := ReadInt()
	k := ReadInt()
	a := make([]int, k+1)
	for i := 1; i <= k; i++ {
		a[i] = ReadInt()
	}
	for i := 1; i <= m; i++ {
		x := ReadInt()
		y := ReadInt()
		add(x, y)
		add(y, x)
	}
	bfs(&dis1,1, n)
	bfs(&dis2,n, n)
	sort.Sort(shuzu(a))
	for i := 1; i <= n; i++ {
		//fmt.Println(dis1[i])
		//fmt.Println(dis2[i])
	}
	for i := 2; i <= k; i++ {
		ans = max(ans, min(dis1[n], 1+dis1[a[i-1]]+dis2[a[i]]))
	}
	fmt.Println(ans)
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