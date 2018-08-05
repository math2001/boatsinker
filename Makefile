runserver: boatsinker
	clear
	./boatsinker

boatsinker: server/**/*.go server/*.go
	cd server; go build -i -o boatsinker
	mv server/boatsinker .

buildclient:
	parcel build src/index.html

watchsrc:
	parcel watch src/index.html --no-hmr

