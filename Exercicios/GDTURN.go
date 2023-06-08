package main

import "fmt"

func main(){
	// your code goes here
	var tc,x,y int;
	fmt.Scanln(&tc);
	for ; tc > 0; tc-- {
	    fmt.Scanf("%d %d",&x,&y)
	    if x + y > 6 {
	        fmt.Println("YES")
	    }else{
	        fmt.Println("NO")
	    }
	}
}
