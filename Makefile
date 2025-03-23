build:
	go build -o bin/antifilter cmd/main.go

# make build gen-youtube
gen-youtube:
	bin/antifilter iplist-youtube/cidr4.txt routes/youtube-ipv4.bat

# make build gen-facebook
gen-facebook:
	bin/antifilter facebook-ip-lists/facebook_ipv4_cidr_blocks.lst routes/facebook-ipv4.bat
