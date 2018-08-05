.SILENT:

runserver:
	make buildserver
	./boatsinker

buildserver:
	cd server; go build -o boatsinker
	mv server/boatsinker .

buildclient:
	parcel build src/index.html

watchsrc:
	parcel watch src/index.html

