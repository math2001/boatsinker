.SILENT:

runserver:
	make buildserver
	./boatkiller

buildserver:
	cd server; go build -o boatkiller
	mv server/boatkiller .

buildclient:
	parcel build src/index.html

watchsrc:
	parcel watch src/index.html

