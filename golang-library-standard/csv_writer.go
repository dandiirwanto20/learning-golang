package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"dandi", "irwanto", "ir"})
	_ = writer.Write([]string{"dian", "joko", "sukarno"})
	_ = writer.Write([]string{"caknan", "deny", "megawati"})

	writer.Flush()
}
