package main

import "strings"

func sort(s string) string {

	k := "lRart"
	kk := []string{}
	for i := 0; i < len(k); i++ {
		for j := 0; j < len(s); j++ {
			if k[i] == s[j] {
				kk = append(kk, string(s[j]))

			}

		}

	}
	m := strings.Join(kk,"")
	return m
}