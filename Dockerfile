ARG ALPINE_VERSION=3.18
ARG GO_VERSION=1.21
ARG GOLANGCI_LINT_VERSION=v1.52.2
ARG MOCKGEN_VERSION=v1.6.0

FROM qmcgaw/binpot:golangci-lint-${GOLANGCI_LINT_VERSION} AS golangci-lint
FROM qmcgaw/binpot:mockgen-${MOCKGEN_VERSION} AS mockgen

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
# Note: findutils needed to have xargs support `-d` flag for mocks stage.
RUN apk --update add git g++ findutils
ENV CGO_ENABLED=0
COPY --from=golangci-lint /bin /go/bin/golangci-lint
COPY --from=mockgen /bin /go/bin/mockgen
WORKDIR /tmp/gobuild
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM --platform=${BUILDPLATFORM} base AS mocks
RUN git init && \
  git config user.email ci@localhost && \
  git config user.name ci && \
  git config core.fileMode false && \
  git add -A && \
  git commit -m "snapshot" && \
  grep -lr -E '^// Code generated by MockGen\. DO NOT EDIT\.$' . | xargs -r -d '\n' rm && \
  go generate -run "mockgen" ./... && \
  git diff --exit-code && \
  rm -rf .git/

FROM base AS test
# Note on the go race detector:
# - we set CGO_ENABLED=1 to have it enabled
# - we installed g++ to support the race detector
ENV CGO_ENABLED=1
ENTRYPOINT go test -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./...

FROM base AS lint
RUN golangci-lint run --timeout=10m


