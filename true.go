package main


func bolle(s string)bool{


	for i := 1; i < len(s); i++ {
		if s[i]!='t'&&s[i]!='R'&&s[i]!='r'&&s[i]!='l'&&s[i]!='a'{
			return false
		}
		
	}
	return true
}