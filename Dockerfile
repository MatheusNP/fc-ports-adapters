FROM golang:1.20

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install go.uber.org/mock/mockgen@latest && \
    go install github.com/spf13/cobra-cli@latest

RUN apt update && apt install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]
