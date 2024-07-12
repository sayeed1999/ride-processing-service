Project Deployment Guide

To set the env's properly, run from bash terminal: -
`export RideProcessingService__Server__Port=8080`
`export RideProcessingService__Server__Host=0.0.0.0`

[Note: While not running with docker compose, omit the first part 'RideProcessingService__'.
While running docker compose, docker will omit the prefix 'RideProcessingService__' for you.]

To run the project directly from terminal: -
Open a terminal from this directory and run `go run ./cmd/app/.`

The api will be running on `localhost:8080`.

---

To manually build the Docker image, run from terminal: -
`docker build -t ride-processing-service .`

To manually run a container for this image, run for terminal: -
`docker run --rm -it -p 7000:8080 ride-processing-service`

The api will be running on `localhost:7000`.

---

To run through Docker Compose file, run from terminal: -
`docker-compose up -d`.
