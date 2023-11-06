package main
import "fmt"
func removeDup(s []string)[]string{
mp:=make(map[string]bool);
var result []string;
for _,str := range s{
    if _,ok:=mp[str]; !ok{
		mp[str]=true;
		result=append(result,str);
	}
}
return result;
}

func main(){

	 a := []string{"fuc","fang","fang"};
	fmt.Println("before",a);
	 res  := removeDup(a);
	fmt.Println("After",res);
}