package main

//翻转字符串中每个单词  思路是 倒序 找到空格的时候 然后 for循环找出这个单词
import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// fmt.Println(len(s))
	s = strings.TrimSpace(s)
	s = " " + s
	res := ""
	flag := len(s)
	for i := len(s) - 1; i >= 0; i -- {
		if (i == 0) || (s[i] == ' ' && s[i - 1] != ' ')  {
			//z找到一个空格
			
			for j := i + 1; j < flag; j ++ {
				if s[j] != ' ' {
					res = res +  string(s[j])
				}
			}
			res = res + " "
			flag = i
		}
	}
	res = res[:len(res) - 1]
	return res
}

func main() {

	str1 := "the sky is blue"
	res := reverseWords(str1) 
	fmt.Println(res)
}