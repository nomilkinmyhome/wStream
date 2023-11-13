## wStream

A simple app for streaming videos.
Just send the GET-request to the ```/stream``` endpoint.
#### And don't forget to copy some video.mp4 into the project directory.

### Run
```go run main.go```

Or with docker:

```docker build -t wStream-app -f infra/Dockerfile .```<br />
```docker run -dp 127.0.0.1:8080:8080 wStream-app```