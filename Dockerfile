FROM alpine
ENV LANGUAGE="en"
WORKDIR ./
COPY . .
RUN apk add --no-cache go
EXPOSE 80/tcp
CMD [ "go", "run", "bot.go" ]