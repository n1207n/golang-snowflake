package main

import snowflake "github.com/n1207n/golang-snowflake-id-generator"

func main() {
	worker := snowflake.NewSnowflakeWorker(1)

	for i := 0; i < 10; i++ {
		println(worker.NextID())
	}
}
