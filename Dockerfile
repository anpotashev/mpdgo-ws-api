FROM scratch
ADD build/mpdapp-linux-amd64 /app

EXPOSE 8080
ENTRYPOINT ["/app"]