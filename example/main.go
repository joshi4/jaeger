package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	zipkinreporter "github.com/openzipkin/zipkin-go/reporter"
	zipkinhttpreporter "github.com/openzipkin/zipkin-go/reporter/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func main() {
	r := NewReporter()
	h := r.InstrumentHandler(http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":10000", h))
}

// Reporter reports metrics.
type Reporter struct {
	reporter zipkinreporter.Reporter
	tracer   *zipkin.Tracer
}

// NewReporter returns a Reporter that reports metrics to url.
func NewReporter() *Reporter {
	host := os.Getenv("COLLECTOR_HOST")
	if len(host) == 0 {
		log.Println(" environment variable COLLECTOR_HOST not set. No traces will be emitted.")
		return nil
	}

	url := fmt.Sprintf("http://%s:9411/api/v2/spans", host)
	reporter := zipkinhttpreporter.NewReporter(url)

	endpoint, err := zipkin.NewEndpoint("example/go", "")
	if err != nil {
		log.Fatalf("unable to create zipkin endpoint: %+v\n", err)
	}

	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create zipkin tracer: %+v\n", err)
	}

	r := &Reporter{
		reporter: reporter,
		tracer:   tracer,
	}
	return r
}

func (r *Reporter) InstrumentHandler(h http.Handler) http.Handler {
	// be defensive
	if r == nil || r.tracer == nil {
		return h
	}

	wrapperFn := zipkinhttp.NewServerMiddleware(
		r.tracer, zipkinhttp.TagResponseSize(true),
	)
	return wrapperFn(h)
}
