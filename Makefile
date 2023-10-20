run-mongo:
	podman run -d -p 27017:27017 --name mongodb mongo

.PHONY: podman-up
pd-up:
	podman-compose up -d

.PHONY: podman-down
pd-down:
	podman-compose down
