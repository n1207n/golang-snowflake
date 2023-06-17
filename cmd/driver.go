package main

import snowflake "github.com/n1207n/golang-snowflake-id-generator"

const ID_SIZE = 1_000_000

func main() {
	worker := snowflake.NewSnowflakeWorker(1)

	for i := 0; i < ID_SIZE; i++ {
		println(worker.NextID())
	}
}
