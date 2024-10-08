HOST=127.0.0.1
build:
	cd ./client && go build -o client && cd ../
http:
	./http_svc/http_svc -host localhost -port 9898 -s_name http_svc_1
up: down docker-rm 
	docker compose up -d
down:
	docker compose down
docker-rm:
	docker rmi traefik-example-http_svc_1 || true
	docker rmi traefik-example-tcp_svc_1 || true
clean:
	rm -rf http_svc/http_svc
logs:
	docker compose logs -f
tcp:
	./client/client -t -host $(HOST) -port 6868
udp:
	./client/client -port 5858
ludp:
	go run ./udp_svc/main.go -port 5858 -s_name udp1
passwd:
	@echo $(shell htpasswd -nB user) | sed -e s/\\$/\\$\\$/g