FROM golang:1.11.4-alpine

LABEL "com.github.action.name"="Go-workflow"
LABEL "com.github.actions.description"="Go test"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

ENV GO111MODULE=on
RUN apk add --no-cache git
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
