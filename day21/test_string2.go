package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo, bar, baz

	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1

	fmt.Println("ba" + strings.Repeat("na", 2)) // banana

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]

}
