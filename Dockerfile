FROM golang:1.20.5-alpine3.18 as build

WORKDIR /src

COPY . ./

RUN go mod download && go build -o ./app ./cmd/server/main.go


FROM golang:1.20.5-alpine3.18 as final

ARG USERNAME=apprunner
# Create the user
RUN addgroup -S  $USERNAME \
    && adduser -S $USERNAME -G $USERNAME

WORKDIR /home/$USERNAME/bin

COPY --from=build /src/app ./
COPY configs configs
RUN chown -R 1000:1000 /home/$USERNAME/

USER $USERNAME
CMD ./app