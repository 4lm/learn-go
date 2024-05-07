# A Tour of Go

## Exercise: Slices

### Link

https://go.dev/tour/moretypes/18

### Solution

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for x := range pic {
		pic[x] = make([]uint8, dx)
		for y := range pic[x] {
			pic[x][y] = uint8(x^y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
```

### Result

![Result Image](./image.png)

## Exercise: Maps

### Link

https://go.dev/tour/moretypes/23

### Solution

```go
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := make(map[string]int)
	for _, f := range fields {
		m[f] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

### Result

```
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
```

## Exercise: Fibonacci closure

### Link

https://go.dev/tour/moretypes/26

### Solution

```go
package main

import "fmt"

func fibonacci() func() int {
	f1, f2 := 1, 0
	return func() int {
		f := f2
		f1, f2 = f1+f, f1
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

### Result

```
0
1
1
2
3
5
8
13
21
34
```

## Exercise: Stringers

### Link

https://go.dev/tour/methods/18

### Solution

```go
package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

### Result

```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

## Exercise: Errors

### Link

https://go.dev/tour/methods/20

### Solution

```go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	g := 1.0 // guess
	prev := 0.0
	for !equal(g, prev) {
		prev = g
		g -= (g*g - x) / (2 * g)
	}
	return g, nil
}

func equal(a, b float64) bool {
	const tolerance = 0.000000000000001
	return math.Abs(a-b) <= tolerance
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

### Result

```
1.414213562373095 <nil>
0 cannot Sqrt negative number: -2
```
