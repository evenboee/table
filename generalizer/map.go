package generalizer

import (
	"fmt"
)

func Map[K comparable, V any](data []map[K]V, explicitOrder []K) ([]string, []map[string]string) {
	rows := make([]map[string]string, len(data))
	headers := make(map[string]struct{})

	for i, row := range data {
		m := make(map[string]string)
		for k, v := range row {
			sK := fmt.Sprintf("%v", k)
			sV := fmt.Sprintf("%v", v)
			m[sK] = sV
			headers[sK] = struct{}{}
		}
		rows[i] = m
	}

	h := make([]string, len(headers)+len(explicitOrder))

	order := make(map[string]int)
	for i, k := range explicitOrder {
		order[fmt.Sprintf("%v", k)] = i
	}

	i := len(explicitOrder)
	for k := range headers {
		if idx, ok := order[k]; ok {
			h[idx] = k
			delete(order, k)
		} else {
			h[i] = k
			i++
		}
	}
	h = h[:i]

	for k, idx := range order {
		h[idx] = k
	}

	return h, rows
}
