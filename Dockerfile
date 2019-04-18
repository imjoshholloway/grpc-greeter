FROM alpine:3.7

# User/Group specification
ARG USER=app-user
ARG GROUP=app-user
ARG UID=510
ARG GID=510
ARG BINARY=app

# Allow us to make HTTPS calls
RUN apk add --update ca-certificates \
    && rm -rf /var/cache/apk/*

# Install the user/group
RUN addgroup -g ${GID} ${GROUP} \
    && adduser -D -u ${UID} -G ${GROUP} ${USER}
USER ${USER}

# Copy binary from build stage
COPY ./${BINARY} /usr/local/bin/app

# Listen on standard ports
EXPOSE 8080

# Specify our launch point
ENTRYPOINT ["/usr/local/bin/app"]