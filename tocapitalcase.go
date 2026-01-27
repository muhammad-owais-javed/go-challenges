package main

func ToCapitalCase(s string) string{

var r = []rune(s)
var tag = true

	for i:=0;i<len(r);i++{

		if r[i]<('a') || r[i]>('z'){
			if r[i]<('A') || r[i]>('Z'){
  			  if r[i]<('0') || r[i]>('9'){
				
				 tag = true
			  }
			}
		}

		if r[i]>='0' && r[i]<='9' {
			tag = false
		}

		if r[i]>=('A') && r[i]<=('Z') {

			if tag == false	{
				r[i] = r[i]+32

			}
			
			tag = false

		}
			
		if r[i]>=('a') && r[i]<=('z') {
			if tag == true	{
				r[i] = r[i]-32
			}
			tag = false
		}

	}

	return string(r)

}
