kafka:
	docker run -p 9092:9092 --network app apache/kafka:3.8.0

d_build:
	docker build -t kafka-projects:local .

d_run:
	docker rm -f messages && docker run --name messages  --network app -p 8083:8083 -v $(pwd)/config-docker.yaml:/app/.env-docker:ro -d kafka-projects:local

d_rm:
	docker rm -f messages