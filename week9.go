package main
import "fmt"
func main(){
	var num[100] int;
	var n,avg,sum int;

	fmt.Print("Enter number of Elemnts:");
	fmt.Scanln(&n);
	for i:=0;i<n;i++{
		fmt.Print("Enter the number :");
		fmt.Scanln(&num[i]);
		sum+=num[i];
	}
	avg = sum/n;
	fmt.Printf("Ther Average of elemnt are %d",avg);
}