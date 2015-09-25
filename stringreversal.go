package main
import "fmt"
func rev( s string) []rune{
	r:= []rune(s)
	for i,j:=0, len(r)-1;i<len(r)/2;i,j=i+1,j-1{
		r[i],r[j]=r[j],r[i]
	}
	fmt.Println(s)		// should print "muthu" and then print result
	return r
}
func main(){
	name:= "muthu"
	new:= rev(name)
   // fmt.Println(new)// this new prints the ascii value of the character and not the character . if we want character string(character).
	fmt.Println("the name in reverse order is ",string(new))
}
