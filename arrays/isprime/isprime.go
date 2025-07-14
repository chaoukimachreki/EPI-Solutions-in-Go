package main

import "fmt"


func generatePrimes(n int) []int {

	//how many prime numbers would be from 0 to n ? 
	//we dont know , so we set the size of the array 
	//to the number of odd numbers between 0 and n since all prime numbers are odd numbers 
	//that we know because we can calculate the last index that i could take 
	//to represent only odd numbers
	//so since we we are given a number n and not a range p-n we will start looking 
	//from 0 to n. 
	//0 is neiher prime not composite, we can eliminate that, 1 is not prime and 2 is prime so we can 
	//safely add it in our array in all cases and therefore always start from 3
	// so to count odd numbers from 3 to n we would do 2*i + 3 which is <= n
	//so the last index would be (n-3/2)
	//the size is last index +1.
	if n <= 2 {return nil}
	if n == 2 {return []int{2}}
	size  := (n-3)/2 +1
	//the idea is  : start with 3 which is prime, all its multiples are not prime
	//so discard those,what ever is left is prime , all its multiples are not prime so 
	//discard those.. and so on. so have two array ready,one to check if it's prime and
	//one to hol the prime numbers consequently

	primes := []int{2}
	isPrime := make([]bool,size)
	for i := range isPrime {isPrime[i]=true}
	
	for i := 0; i < size; i++ {

		p := 2*i + 3
		if isPrime[i] {primes = append(primes,p)}
		for j := i+p; j <= size; j+=p {isPrime[j]=false}
	}
	return primes


}

func main(){
	
	n := 2
	primes := generatePrimes(n)
	fmt.Printf("here is the list of prime numbers that are less than %d\n : %v\n",n,primes)


}
