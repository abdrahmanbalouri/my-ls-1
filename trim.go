package main


func trim(s string)string{

	for i:=len(s)-1;i>=0;i--{

		if s[i]=='/'{
			s=s[i+1:]
			break
		}
	}

  
  return s

}