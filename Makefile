build:
	go build -o bin/antifilter cmd/main.go

# make build gen-youtube
gen-youtube:
	bin/antifilter iplist-youtube/cidr4.txt routes/youtube-ipv4.bat

# make build gen-facebook
gen-facebook:
	bin/antifilter facebook-ip-lists/facebook_ipv4_cidr_blocks.lst routes/facebook-ipv4.bat

# make build gen-chatgpt
gen-chatgpt:
	bin/antifilter cidr/chatgpt_cidr4.txt routes/chatgpt-ipv4.bat

# make build gen-medium
gen-medium:
	bin/antifilter cidr/medium_cidr4.txt routes/medium-ipv4.bat

gen-all: gen-youtube
gen-all: gen-facebook
gen-all: gen-chatgpt
gen-all: gen-medium
