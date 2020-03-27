FROM jaegertracing/jaeger-collector:1.17.0

# ES_SERVER_HOST is set in render.yml
ENV SPAN_STORAGE_TYPE=elasticsearch
ARG ES_SERVER_HOST
ENV ES_SERVER_URLS="http://${ES_SERVER_HOST}:9200"
