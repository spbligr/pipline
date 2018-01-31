package main

import (
	"fmt"
	//"sync"
	//"log"
	"sync"
)

func ExecutePipeline(freeFlowJobs... job)  {

	in := make(chan interface{})
	out := make(chan interface{})
	//wgOut := make(chan interface{})
	//wgIn := make(chan int)


	wg := &sync.WaitGroup{}

	for _, job := range freeFlowJobs {

		wg.Add(1)

		go func(jobFunc func(in chan interface{}, out chan interface{}), in chan interface{}, out chan interface{}) {
			go jobFunc(in, out)
			select {
			case <- in:
				fmt.Println("wgDone")
				wg.Done()
			}
		}(job, in, out)

	}

	wg.Wait()

}

var SingleHash = func(in, out chan interface{})  {
	for value := range out {

		value = fmt.Sprintf("%s", value)

		hashData := make(map[string]string)
		result := DataSignerCrc32(hashData["data"]) + "~" + DataSignerCrc32(DataSignerMd5(hashData["data"]))

		fmt.Println("SingleHash data= , resul = ", value, result)

		in <- result
	}
}

var MultiHash = func(in, out chan interface{}) {
	for value := range in {
		fmt.Println("MultiHash", value)
	}

}

var CombineResults = func(in, out chan interface{}) {

}