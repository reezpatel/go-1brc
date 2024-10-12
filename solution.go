package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Station struct {
	l, h, s float64
	c int
}

func solve(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var items = make(map[string]Station)

	len := 0

    for scanner.Scan() {
		s := scanner.Text()

		parts := strings.Split(s, ";")

		n, err := strconv.ParseFloat(parts[1], 32)

		if err != nil {
			panic(err)
		}

		e, ok := items[parts[0]]

		if !ok {
			len += 1
			items[parts[0]] = Station{
				l: n,
				h: n,
				s: n,
				c: 1,
			}
		} else {
			e.l = min(e.l, n)
			e.h = max(e.h, n)
			e.s += n
			e.c += 1

			items[parts[0]] = e
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	var keys = make([]string, len)

	i := 0
	for k := range items {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	arr := make([]string, len)

	i = 0
	for _, k := range keys {
		arr[i] = fmt.Sprintf("%s=%.1f/%.1f/%.1f", k, items[k].l, items[k].s / float64(items[k].c), items[k].h)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(arr, ", "))
}
