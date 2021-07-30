
package main

import (
	"fmt"
	"golang.org/x/net/idna"
)

var p *idna.Profile

func main() {
	emojis := []string{"🐛🐛", "🐈🎊", "🤣😻", "😻🍺", "😱🐈", "🐈😇", "🤣🤢", "😻🌮", "😱🤣", "🐈😘", "🤣🐛", "😻😇", "😱😻", "🐈🐛", "🤣🐛", "😻🐛", "😱😱"}
	p = idna.New()

	for i, s := range emojis {
		fmt.Print(emojis[i], " ")
		fmt.Println(p.ToASCII(s))
	}
}
