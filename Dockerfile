FROM registry.suse.com/bci/golang:1.19-18.38 as base
WORKDIR /app
COPY . .
RUN go build -ldflags "-s -w" -o fleet-ui

FROM registry.suse.com/bci/bci-base:15.4.27.14.20
WORKDIR /app
COPY --from=base /app/fleet-ui .
CMD ["/app/fleet-ui"]
