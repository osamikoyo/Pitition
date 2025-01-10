FROM ubuntu:latest
LABEL authors="osami"

ENTRYPOINT ["top", "-b"]