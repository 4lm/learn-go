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
