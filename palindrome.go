package main
import "fmt"
func rev( s string) string{
	r:= []rune(s)
	for i,j:=0, len(r)-1;i<len(r)/2;i,j=i+1,j-1{
		r[i],r[j]=r[j],r[i]
	}
	fmt.Println(s)		// should print "muthu" and then print result
	return string(r)
}
func main(){
	var name string
	fmt.Println("Enter the string you want to reverse")
	fmt.Scanf("%s",&name)
	fmt.Println("the name is ",name)
	new:= rev(name)
	fmt.Println("The reversal value of string is", new)// this new prints the ascii value of the character and not the character . if we want character string(character).
    if (name == new){
		fmt.Println("the string is a palindrome")
	}else{
		fmt.Println("the string is not a palindrome")
}
}
