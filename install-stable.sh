#!/bin/sh
# shellcheck shell=dash


set -u

GITHUB_ORG=CoopHive
GITHUB_REPO=hive

PRE_RELEASE=${PRE_RELEASE:=false}
COOPHIVE_HTTP_REQUEST_CLI=${COOPHIVE_HTTP_REQUEST_CLI:="curl"}

version=${version:="v0.5.4"} #change arc back to amd64 once updating to latest version

detect_os_info() {
  OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "x86_64"; else print "unsupported_arch"}') && export OSARCH
  OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "Darwin"; else if ($1 == "Linux") print "Linux"; else print "unsupported_os"}') && export OSNAME

  if [ "$OSNAME" = "unsupported_os" ] || [ "$OSARCH" = "unsupported_arch" ]; then
    echo "Unsupported OS or architecture"
    exit 1
  fi
}

install_hive() {
  echo "installing hive:$version"
  rurl="https://github.com/CoopHive/hive/releases/download/$version/hive-$OSNAME-$OSARCH"
  echo "Release url is $rurl"
  curl -sSL -o hive "$rurl"
  chmod +x hive
  ./hive version
  sudo cp hive /usr/local/bin/
}

install_bacalhau() {
  curl -sL https://get.bacalhau.org/install.sh | bash
}

main() {
  detect_os_info


  if [ "$#" -eq 0 ]; then
    echo "Usage: $0 [all|bacalhau|hive]"
    exit 1
  fi


  if [ "$1" = "all" ]; then
    install_hive
    install_bacalhau
  elif [ "$1" = "bacalhau" ]; then
    install_bacalhau
  elif [ "$1" = "hive" ]; then
    install_hive
  else
    echo "Usage: $0 [all|bacalhau|hive]"
    exit 1
  fi
}

main "$@" || exit 1
