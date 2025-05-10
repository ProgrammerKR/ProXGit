# Build stage
FROM docker.io/library/golang:1.24-alpine3.21 AS build-env

ARG GOPROXY
ENV GOPROXY=${GOPROXY:-direct}

ARG GITEA_VERSION
ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS="bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

# Build deps
RUN apk --no-cache add \
    build-base \
    git \
    nodejs \
    npm \
    && rm -rf /var/cache/apk/*

# Setup repo
<<<<<<< HEAD
COPY . ${GOPATH}/src/code.gitea.io/gitea
WORKDIR ${GOPATH}/src/code.gitea.io/gitea
=======
COPY . ${GOPATH}/src/code.proxgit.io/proxgit
WORKDIR ${GOPATH}/src/code.proxgit.io/proxgit
>>>>>>> master

# Checkout version if set
RUN if [ -n "${GITEA_VERSION}" ]; then git checkout "${GITEA_VERSION}"; fi \
 && make clean-all build

# Begin env-to-ini build
RUN go build contrib/environment-to-ini/environment-to-ini.go

# Copy local files
COPY docker/root /tmp/local

# Set permissions
RUN chmod 755 /tmp/local/usr/bin/entrypoint \
<<<<<<< HEAD
              /tmp/local/usr/local/bin/gitea \
              /tmp/local/etc/s6/gitea/* \
              /tmp/local/etc/s6/openssh/* \
              /tmp/local/etc/s6/.s6-svscan/* \
              /go/src/code.gitea.io/gitea/gitea \
              /go/src/code.gitea.io/gitea/environment-to-ini
RUN chmod 644 /go/src/code.gitea.io/gitea/contrib/autocompletion/bash_autocomplete

FROM docker.io/library/alpine:3.21
LABEL maintainer="maintainers@gitea.io"
=======
              /tmp/local/usr/local/bin/proxgit \
              /tmp/local/etc/s6/proxgit/* \
              /tmp/local/etc/s6/openssh/* \
              /tmp/local/etc/s6/.s6-svscan/* \
              /go/src/code.proxgit.io/proxgit/proxgit \
              /go/src/code.proxgit.io/proxgit/environment-to-ini
RUN chmod 644 /go/src/code.proxgit.io/proxgit/contrib/autocompletion/bash_autocomplete

FROM docker.io/library/alpine:3.21
LABEL maintainer="maintainers@proxgit.io"
>>>>>>> master

EXPOSE 22 3000

RUN apk --no-cache add \
    bash \
    ca-certificates \
    curl \
    gettext \
    git \
    linux-pam \
    openssh \
    s6 \
    sqlite \
    su-exec \
    gnupg \
    && rm -rf /var/cache/apk/*

RUN addgroup \
    -S -g 1000 \
    git && \
  adduser \
    -S -H -D \
    -h /data/git \
    -s /bin/bash \
    -u 1000 \
    -G git \
    git && \
  echo "git:*" | chpasswd -e

ENV USER=git
<<<<<<< HEAD
ENV GITEA_CUSTOM=/data/gitea
=======
ENV GITEA_CUSTOM=/data/proxgit
>>>>>>> master

VOLUME ["/data"]

ENTRYPOINT ["/usr/bin/entrypoint"]
CMD ["/usr/bin/s6-svscan", "/etc/s6"]

COPY --from=build-env /tmp/local /
<<<<<<< HEAD
COPY --from=build-env /go/src/code.gitea.io/gitea/gitea /app/gitea/gitea
COPY --from=build-env /go/src/code.gitea.io/gitea/environment-to-ini /usr/local/bin/environment-to-ini
COPY --from=build-env /go/src/code.gitea.io/gitea/contrib/autocompletion/bash_autocomplete /etc/profile.d/gitea_bash_autocomplete.sh
=======
COPY --from=build-env /go/src/code.proxgit.io/proxgit/proxgit /app/proxgit/proxgit
COPY --from=build-env /go/src/code.proxgit.io/proxgit/environment-to-ini /usr/local/bin/environment-to-ini
COPY --from=build-env /go/src/code.proxgit.io/proxgit/contrib/autocompletion/bash_autocomplete /etc/profile.d/proxgit_bash_autocomplete.sh
>>>>>>> master
