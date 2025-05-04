FROM golang:latest AS builder

WORKDIR /src
COPY . .
RUN rm -rf _deploy/index.html _gh-pages
RUN go build
RUN ./resume

# ---
FROM halverneus/static-file-server:latest
COPY --from=builder /src/_deploy /public

ENV URL_PREFIX=/resume
ENV FOLDER=/public

EXPOSE 8080
