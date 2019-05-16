FROM alpine

ENV DEFAULT_PATH mypath

WORKDIR /usr/share/docker

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

RUN apk update \
    && apk add curl \
    && LAST_VERSION=$(curl -sqLk https://api.github.com/repos/leocomelli/simple-api/releases/latest | grep tag_name | head -n 1 | cut -d '"' -f 4) \
    && curl -sLk https://github.com/leocomelli/simple-api/releases/download/${LAST_VERSION}/simple-api_linux-amd64 -o simple-api_linux-amd64 \
    && chmod +x simple-api_linux-amd64 /usr/local/bin/entrypoint.sh

ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]