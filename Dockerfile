FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-web-app"]
COPY ./bin/ /