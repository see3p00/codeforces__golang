package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
 
var ch [1000005][26] int
var id [1000005] int
var a [1000005]int
var dp [1000005]int
var vis[1000005]bool
var tot int
 
func dfs1(u int) {
	if vis[u] ==true {
		tot++
		id[u] = tot
	}
	for i := 0; i < 26; i++ {
		if ch[u][i] != 0 {
			dfs1(ch[u][i])
		}
	}
}
func dfs2(u, w int) {
	val:=max(w,tot-dp[u]);
	if vis[u]==true{
		tot++
	}
	for i:=0;i<26;i++{
		v:=ch[u][i]
		if v==0 {
			continue;
		}
		dp[v]=dp[u]+1
		if vis[v]==true{
			dp[v]=min(dp[v],id[v]-val)
		}
		dfs2(v,val)
	}
}
func main() {
	n := ReadInt()
	for i := 1; i <= n; i++ {
		x:=ReadInt()
		s:=ReadString()
		ch[x][s[0]-'a']=i
	}
	q:=ReadInt()
	for i:=1;i<=q;i++{
		a[i]=ReadInt()
		vis[a[i]]=true
	}
	tot=0;dfs1(0)
	tot=0;dfs2(0,0)
	for i:=1;i<=q;i++{
		fmt.Printf("%d ",dp[a[i]])
	}
}
 
////////////////////////////////////
 
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
func ReadInt64() int64 {
	intstr := ReadString()
	i, _ := strconv.ParseInt(intstr, 10, 64)
	return int64(i)
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