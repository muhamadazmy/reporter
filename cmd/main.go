package main

import (
	"fmt"

	"github.com/rivine/reporter"
)

func main() {
	// exp, err := reporter.NewExplorer("http://localhost:23110")
	// if err != nil {
	// 	panic(err)
	// }

	cl, err := reporter.NewInfluxDB("http://localhost:8086/rivine")
	if err != nil {
		panic(err)
	}

	query := reporter.NewInfluxQuery(cl)
	fmt.Println(query.TotalTokens())
	fmt.Println(query.TransactedToken(reporter.LastDay))
	// rep, err := reporter.NewInfluxRecorder(cl, 200, 30*time.Second)
	// if err != nil {
	// 	panic(err)
	// }

	// defer rep.Close()

	// scanner := exp.Scan(0)

	// for blk := range scanner.Scan(context.Background()) {
	// 	if err := rep.Record(blk); err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println(scanner.Err())
}
