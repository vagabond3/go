package main
import(
	"fmt"
	"sort"
)
func main(){
     num :=[]int {1,2,3,4};
     str :=[]string{"z","a","c"};
  
	 sort.Sort(sort.Reverse(sort.IntSlice(num)));
	 sort.Sort(sort.Reverse(sort.StringSlice(str)));
	
	
	fmt.Println(num);
	fmt.Println(str);
}
	

