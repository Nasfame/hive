default: go build 

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
	" .
	./hive version