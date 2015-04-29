package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/daddye/json-bench"
)

func main() {
	log.SetFlags(0)
	f, err := os.OpenFile("./testdata/data.json", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	j := json.NewEncoder(f)

	for i := 0; i < json_bench.Messages; i++ {
		j.Encode(json_bench.Message{"message-" + strconv.Itoa(i), json_bench.Status(i)})
	}
}
