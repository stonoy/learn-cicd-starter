FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD learn-cicd-starter /usr/bin/notely

CMD ["notely"]
