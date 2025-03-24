build:
	go build -o bin/antifilter cmd/main.go

# make build gen-youtube
gen-youtube:
	echo > cidr4.full.ignore.txt
	cat vendor/iplist-youtube/cidr4.txt >> cidr4.full.ignore.txt
	echo "\n" >> cidr4.full.ignore.txt
	cat cidr/youtube_cidr4.txt >> cidr4.full.ignore.txt
	bin/antifilter cidr4.full.ignore.txt routes/youtube-ipv4.bat

# make build gen-facebook
gen-facebook:
	bin/antifilter vendor/facebook-ip-lists/facebook_ipv4_cidr_blocks.lst routes/facebook-ipv4.bat

# make build gen-chatgpt
gen-chatgpt:
	bin/antifilter cidr/chatgpt_cidr4.txt routes/chatgpt-ipv4.bat

# make build gen-medium
gen-medium:
	bin/antifilter cidr/medium_cidr4.txt routes/medium-ipv4.bat

# make build gen-all
gen-all: gen-youtube
gen-all: gen-facebook
gen-all: gen-chatgpt
gen-all: gen-medium
