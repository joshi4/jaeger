# Jaeger On Prem Deployment 

Use the button below to Deploy Jaeger with one click.

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

This deploy has the following components: 
* Jaeger Collector: Receives Traces from a service and writes to Elasticsearch. 
* Jaeger UI: GUI to query Elasticsearch and fetch traces.
* Jaeger Storage: An Elasticsearch node.
* Jager Example: App to emit traces.

## Jaeger Collector 
The collector is a [private service](https://render.com/docs/private-services) that is only accessible from other services running in your render account. It is configured to accept [Zipkin](https://zipkin.io/pages/tracers_instrumentation.html) compatible traces.

By default, the collector will [sample](https://www.jaegertracing.io/docs/1.17/sampling/#collector-sampling-configuration) `0.01%` of spans it receives. 
### Sending Traces to the Collector

Users can use any zipkin compatible sdk to send traces to the Jaegor Collector. Clients(your service) should send their traces to `$COLLECTOR_HOST:9411/api/v2/spans`.

`COLLECTOR_HOST` will be of the form: `jaeger-collector:9411`. The exact hostname can be viewed from the service dashboard on render.

## Jaeger UI 
The Jaeger UI service is a public web service. Users can visit the site and query any trace stored in elasticsearch.

## Jaeger Storage 
Jaeger Storage is [private service](https://render.com/docs/private-services) that runs a single elasticsearch node. Read the elasticsearch [README](./storage/README.md) for more details.


## Jaeger Example 
A simple Go server that answers render healthchecks and emits traces. It acts as an end to end test and can be deleted after verifying that all jaeger components were deployed successfully.
