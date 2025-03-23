build:
	go build -o bin/antifilter cmd/main.go

# make build gen-youtube
gen-youtube:
	bin/antifilter iplist-youtube/cidr4.txt routes/youtube-ipv4.bat
