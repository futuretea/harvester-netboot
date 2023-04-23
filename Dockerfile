FROM golang:1.18.7 as builder
COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 go build -o /app .

FROM registry.suse.com/bci/bci-base:15.4
RUN zypper -n rm container-suseconnect && \
    zypper -n install curl && \
    zypper -n clean -a && rm -rf /tmp/* /var/tmp/* /usr/share/doc/packages/*

RUN curl -sLf https://github.com/krallin/tini/releases/download/v0.19.0/tini > /usr/bin/tini && chmod +x /usr/bin/tini

COPY --from=builder /app /usr/bin/
COPY entrypoint.sh  /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]