package json_bench

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/ugorji/go/codec"
)

func BenchmarkCodec(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	var h codec.JsonHandle
	dec := codec.NewDecoder(f, &h)

	b.ResetTimer()
	var i int
	for ; ; i++ {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
	}

	if i != Messages {
		b.Fatalf("expected %d messages got %d", Messages, i)
	}
}
func BenchmarkCodecBuffer(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	buf := bufio.NewReader(f)

	var h codec.JsonHandle
	dec := codec.NewDecoder(buf, &h)

	b.ResetTimer()
	var i int
	for ; ; i++ {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
	}

	if i != Messages {
		b.Fatalf("expected %d messages got %d", Messages, i)
	}
}

func BenchmarkStdLib(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	dec := json.NewDecoder(f)

	b.ResetTimer()
	var i int
	for ; ; i++ {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
	}

	if i != Messages {
		b.Fatalf("expected %d messages got %d", Messages, i)
	}
}
