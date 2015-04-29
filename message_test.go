package json_bench

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/ugorji/go/codec"
)

func printStats(t metrics.Timer) {
	fmt.Printf("%.2f ops - 99%%: %s - mean: %s\n",
		t.RateMean(),
		time.Duration(t.Percentile(0.99)),
		time.Duration(t.Mean()),
	)
}

func BenchmarkCodec(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	timer := metrics.NewTimer()

	var h codec.JsonHandle
	dec := codec.NewDecoder(f, &h)

	b.ResetTimer()
	for {
		start := time.Now()
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
		timer.UpdateSince(start)
	}

	if timer.Count() != Messages {
		b.Fatalf("expected %d messages got %d", Messages, timer.Count())
	}

	printStats(timer)
}

func BenchmarkCodecBytes(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	timer := metrics.NewTimer()

	b.ResetTimer()
	var h codec.JsonHandle
	dec := codec.NewDecoderBytes(data, &h)

	for {
		start := time.Now()
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
		timer.UpdateSince(start)
	}

	if timer.Count() != Messages {
		b.Fatalf("expected %d messages got %d", Messages, timer.Count())
	}

	printStats(timer)
}

func BenchmarkCodecBuffer(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	timer := metrics.NewTimer()
	buf := bufio.NewReader(f)

	var h codec.JsonHandle
	dec := codec.NewDecoder(buf, &h)

	b.ResetTimer()
	for {
		start := time.Now()
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
		timer.UpdateSince(start)
	}

	if timer.Count() != Messages {
		b.Fatalf("expected %d messages got %d", Messages, timer.Count())
	}

	printStats(timer)
}

func BenchmarkStdLib(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	timer := metrics.NewTimer()
	dec := json.NewDecoder(f)

	b.ResetTimer()
	for {
		start := time.Now()
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			b.Fatal(err)
		}
		timer.UpdateSince(start)
	}

	if timer.Count() != Messages {
		b.Fatalf("expected %d messages got %d", Messages, timer.Count())
	}

	printStats(timer)
}

func BenchmarkStdLibUnmarshal(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("./testdata/data.json")
	if err != nil {
		b.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	timer := metrics.NewTimer()

	b.ResetTimer()
	for scanner.Scan() {
		timer.Time(func() {
			var m Message
			if err := json.Unmarshal(scanner.Bytes(), &m); err != nil {
				b.Fatal(err)
			}
		})
	}

	if timer.Count() != Messages {
		b.Fatalf("expected %d messages got %d", Messages, timer.Count())
	}

	printStats(timer)
}
