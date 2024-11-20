FROM ubuntu:latest
LABEL authors="noskov"

ENTRYPOINT ["top", "-b"]