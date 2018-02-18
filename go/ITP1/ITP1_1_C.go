package main

import(
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	fmt.Println(a*b,2*(a+b))
}