FROM debian:11-slim
RUN apt-get update \
  && apt-get install -y --no-install-recommends curl jq ca-certificates
RUN curl -sLJO -H 'Accept: application/octet-stream' \
  "https://api.github.com/repos/ookkoouu/ExpandMediaBot/releases/assets/$(\
  curl -sL https://api.github.com/repos/ookkoouu/ExpandMediaBot/releases/latest \
  |jq '.assets[] | select(.name=="app_arm") | .id')" \
  && chmod +x app_arm

ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
CMD [ "./app_arm" ]