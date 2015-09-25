package main
import"fmt"
func con(name1,name2 string)string{
	r1 :=[]rune(name1)
	r2 :=[]rune(name2)
	r :=[]rune(name1+name2)
	for i:=0;i<len(name1);i++{
		r[i] = r1[i]
	}
	for i,j:= len(name1),0;j<len(name2);i,j=i+1,j+1{
		r[i]=r2[j]
	}
	return string(r)
}
func main(){
	var name1,name2 string
	fmt.Println("Enter the first name")
	fmt.Scanf("%s\t",&name1)
	fmt.Println("the first name is ",name1)
	fmt.Println("Enter the last name")
	fmt.Scanf("%s\t",&name2)
	fmt.Println("the last name is ",name2)
	new:= con(name1,name2)
	fmt.Println("the full name is",new)// this new prints the ascii value of the character and not the character . if we want character string(character).

}
