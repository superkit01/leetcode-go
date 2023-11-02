package easy

import "golang.org/x/exp/slices"

func CountPoints(rings string) int {
	container := make(map[byte][]byte, 0)

	for i := 1; i < len(rings); i += 2 {
		if _, ok := container[rings[i]]; !ok {
			container[rings[i]] = make([]byte, 0)
		}
		container[rings[i]] = append(container[rings[i]], rings[i-1])
	}

	result := 0

	for _, v := range container {
		if len(v) < 3 {
			continue
		}

		if !slices.Contains(v, 'B') {
			continue
		}

		if !slices.Contains(v, 'G') {
			continue
		}

		if !slices.Contains(v, 'R') {
			continue
		}
		result++
	}

	return result

}
