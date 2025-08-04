package iteration

import "strings"

func Repeat(word string, count int) string {
	var  builder strings.Builder
	for i:=0;i<count;i++ {
		 builder.WriteString(word)
	}
	return builder.String()
}