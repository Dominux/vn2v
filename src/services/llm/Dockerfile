# FROM golang:1.24rc1-bookworm AS build
# RUN mkdir /app
# ADD . /app/
# WORKDIR /app
# RUN go mod download
# RUN go build -o /main ./main.go
# CMD ["/main"]

FROM ollama/ollama:0.9.7-rc0-rocm

RUN apt update && apt install -y curl git-lfs

COPY ./run.sh /
COPY ./Modelfile /

ENTRYPOINT [ "sh" ]
CMD [ "/run.sh" ]
