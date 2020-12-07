FROM scratch

COPY bin/testing-github-action /usr/bin/testing-github-action

EXPOSE 3000
ENTRYPOINT ["/usr/bin/testing-github-action"]
