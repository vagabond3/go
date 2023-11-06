package main
import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("CMRIT","Fu"));//checks substring
	fmt.Println(strings.ContainsAny("cmr","f"));//cehcks wherthe char is prestnt or not
	fmt.Println(strings.Count("cmr","F"));//counts how mny
	fmt.Println(strings.EqualFold("cmr","FU"));
}
