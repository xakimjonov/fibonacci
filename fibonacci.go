package main

import (
	"fmt"
	"time"
)

type calculation struct{
	number, result int
}
func main(){
	const jobsCount = 45
	jobs := make(chan calculation, jobsCount)
	results := make(chan calculation, jobsCount)
    start := time.Now()

	for w:=1; w <= 5; w++ {

		go worker(w, jobs, results)
	}
    
	for j := 1; j<= jobsCount; j++{
		calc := calculation{j,0}
		jobs <- calc
	}
	close(jobs)

	for a:= 1; a<= jobsCount; a++{
		result := <- results
		fmt.Println("Fib(",result.number,")=",result.result)
	}
		close(results)
		duration := time.Since(start)
		fmt.Println("Done! Spend time:",duration.Milliseconds())
}

func worker(id int, jobs <-chan calculation, results chan<- calculation) {
	for j:= range jobs{
		fmt.Println(id, "- worker", j.number, "start this job")
		j.result = calculateFibonacci(j.number)
		results <- j
		fmt.Println(id, "-worker", j.number, "finished this job")
	}
}

func calculateFibonacci(n int) int {
	if n<= 1{
		return n
	}
	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
	
	
}
