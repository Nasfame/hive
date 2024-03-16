set shell := ["sh", "-c"]
set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]
#set allow-duplicate-recipe
set positional-arguments
set dotenv-load
set export


default:
    go build

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
	" .
	./hive version

echo:
    @echo just

setup-geth:
    echo "Setup geth"
