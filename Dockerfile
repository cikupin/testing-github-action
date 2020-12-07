FROM scratch

COPY bin/testing-github-action /go/bin/testing-github-action

EXPOSE 3000
ENTRYPOINT ["/go/bin/testing-github-action"]