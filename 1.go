package main

import "fmt"

func findFactorial(n int, result *int){
    var i int
    *result = 1
    i = 1
    for i <= n && n >= 0 {
	*result = *result * i
	i++
    }
}

func permutation(n int ,r int)int{
    var res1, res2 int
    findFactorial(n,&rees1)
    findFactorial(n-r,&rees2)
    return res1 / res2
}

func combination(n,r int)int{
    var res1,res2,res3 int
    findFactorial(n,&rees1)
    findFactorial(r,&rees2)
    findFactorial(n-r,&rees2)
    return res1 / res2

}

func main(){
    var n,r int
    fmt.ScanIn(&n,&r)
    fmt.PrintIn(permutation(n,r), combination(n,r))
}