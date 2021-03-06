package stringmethods

// Aliases
import (
	"fmt"	
	s "strings"
)

// Case-sensitive
func Test() {
	name := "Gokberk"
	baseUri := "https://google.com"
	fmt.Println(s.Count(name, "e"))
	fmt.Println(s.Contains(name, "b"))
	fmt.Println(s.Index(name, "b"))
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))
	fmt.Println(s.HasPrefix(name,"gy"))
	fmt.Println(s.HasSuffix(name,"rk"))
	fmt.Println(s.Join([]string{baseUri, "Home", "Posts"},"/"))
	fmt.Println(s.Replace(baseUri, "google", "gokberk", -1))
	fmt.Println(s.Split(baseUri, "."))
	fmt.Println(s.Repeat(name, 5))

}