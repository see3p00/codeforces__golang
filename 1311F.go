package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

var ans int64
var c [200005][2]int64
var n int

type node struct {
	x int
	v int
}

func lowbit(x int) int {
	return x & (-x)
}
func update(i, val int) {
	for i <= n {
		c[i][0] ++
		c[i][1] += int64(val)
		i += lowbit(i)
	}
}
func query(i, k int) int64 {
	var res int64 = 0
	for i > 0 {
		res += c[i][k]
		i -= lowbit(i)
	}
	return res
}
func unique_sort(slice []int) []int {
	res := make([]int, 0)
	mp := make(map[int]bool, len(slice))
	for _, e := range slice {
		if mp[e] == false {
			mp[e] = true
			res = append(res, e)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}
func lower_bound(st int, ed int, key int, num []int) int {
	left, right := 1, ed-1
	for ; left < right; {
		mid := (right + left) / 2
		if num[mid] >= key {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == ed {
		return ed
	} else {
		return right
	}
}
func main() {
	n = ReadInt()
	a := make([]node, n+1)
	var b []int
	a[0].x = -1e9
	a[0].v = -1e9
	for i := 1; i <= n; i++ {
		a[i].x = ReadInt()
	}
	b = append(b, -1e9)
	for i := 1; i <= n; i++ {
		a[i].v = ReadInt()
		b = append(b, a[i].v)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x;
	})
	b = unique_sort(b)
	for i := 1; i <= n; i++ {
		tmp := lower_bound(1, len(b), a[i].v, b)
		ans += int64(a[i].x) * query(tmp, 0) - query(tmp, 1)
		update(tmp, a[i].x)
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