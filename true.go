package main


func bolle(s []string)bool{


	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s[i]); j++ {
			if s[i][j]!='t'&&s[i][j]!='R'&&s[i][j]!='r'&&s[i][j]!='l'&&s[i][j]!='a'{
			return false
		}
		}
		
		
	}
	return true
}