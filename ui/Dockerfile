FROM jaegertracing/jaeger-query:1.17.0

# declared in render.yml for service name: jaeger-ui
ENV SPAN_STORAGE_TYPE=elasticsearch
ARG ES_SERVER_HOST 
ENV ES_SERVER_URLS="http://${ES_SERVER_HOST}:9200"
