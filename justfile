default: go build 

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .
	./hive version