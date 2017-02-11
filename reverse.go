package main

func ReverseString(str string) string {
	if len(str) <= 1 {
		return str
	} else {
		return ReverseString(str[1:]) + str[:1]
	}
}

