package main
//strings paketi ile string birleştirme
//performanslı string birleştirme işlemi
//en performanslı birleştirme string builder ile yapılan
import(
	"fmt"
	"strings"
)
func main(){
	var builder strings.Builder
	for i:=0;i<100 ;i++ {
		builder.WriteString("-deneme mesajı")
	}
	fmt.Println(builder.String())
}
