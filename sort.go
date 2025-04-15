package main

import "strings"

func sort(s []string) string {

	k := "Ratrl"
	kk := []string{}

	for i := 0; i <len(k) ; i++ {
		for j := 0; j < len(s); j++ {
			for c:= 1; c < len(s[j]); c++ {
				if s[j][c]== k[i]{
					kk = append(kk, string(s[j][c]))


				}
				
			}
			
		}
		
	}
	m:=strings.Join(kk,"")
	
	return m
}