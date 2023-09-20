validate_version:
ifndef VERSION
	$(error VERSION is undefined)
endif

release: validate_version
	# linux
	GOOS=linux go build -ldflags "-X main.version=${VERSION}" -o serve ;\
	tar -zcvf ./releases/serve_${VERSION}_linux.tar.gz ./serve ;\

	# macos
	GOOS=darwin go build -ldflags "-X main.version=${VERSION}" -o serve ;\
	tar -zcvf ./releases/serve_${VERSION}_macos.tar.gz ./serve ;\

	# windows
	GOOS=windows go build -ldflags "-X main.version=${VERSION}" -o serve ;\
	tar -zcvf ./releases/serve_${VERSION}_windows.tar.gz ./serve ;\

	rm ./serve