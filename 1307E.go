package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
 
var cnt [5005][5005] int
var a [5005] int
var tot [5005] int
var l [5005] int
var maxx, ans, n, m int
var mod int64 = 1e9+7
func main() {
	n = ReadInt()
	m = ReadInt()
	for i := 1; i <= n; i++ {
		a[i] = ReadInt()
	}
	for i := 1; i <= m; i++ {
		u := ReadInt()
		v := ReadInt()
		cnt[u][v]++
	}
	for i := 1; i <= n; i++ {
		tot[a[i]]++
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			cnt[i][j] += cnt[i][j-1]
		}
	}
	res:=0
	var sum int64 =1
	for i:=1;i<=n;i++{
		u:=cnt[i][tot[i]]
		if u>0 {
			res++
			sum=sum*int64(u)%mod
		}
	}
	maxx=res
	ans=int(sum)
	for i:=1;i<=n;i++ {
		l[a[i]]++
		if cnt[a[i]][l[a[i]]]==cnt[a[i]][l[a[i]]-1] {
			continue
		}
		res:=1
		var sum int64=1
		u:=cnt[a[i]][tot[a[i]]-l[a[i]]]
		if tot[a[i]]-l[a[i]]>=l[a[i]] {
			u--
		}
		if u>0 {
			res++
			sum = sum * int64(u) % mod
		}
		for j:=1;j<=n;j++ {
			if j==a[i] {
				continue
			}
			p:=l[j]
			q:=tot[j]-l[j]
			u:=cnt[j][p]*cnt[j][q]-cnt[j][min(p,q)]
			if u>0 {
				res+=2
				sum=sum*int64(u)%mod
			} else {
				u=cnt[j][p]+cnt[j][q]
				if u>0 {
					res++
					sum=sum*int64(u)%mod
				}
			}
		}
		if res>maxx {
			maxx=res
			ans=int(sum)
		} else if res==maxx {
			ans=(ans+int(sum))%int(mod)
		}
	}
	fmt.Println(maxx,ans)
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
