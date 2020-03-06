//https://codeforces.ml/contest/1307/problem/A

package main
 
import "fmt"
 
func main() {
	var a [105]int
	var n, d, t, i, cas int
	fmt.Scan(&t)
	for cas = 0; cas < t; cas++ {
		fmt.Scan(&n, &d);
		for i = 1; i <= n; i++ {
			fmt.Scan(&a[i]);
		}
		for i = 2; i <= n; i++ {
			a[1] += min(a[i], d/(i-1))
			d -= min(a[i], d/(i-1)) * (i - 1)
 
		}
		fmt.Println(a[1])
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}