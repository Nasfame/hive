#!/bin/sh
# shellcheck shell=dash

set -u

GITHUB_ORG=CoopHive
GITHUB_REPO=hive

PRE_RELEASE=${PRE_RELEASE:=false}
COOPHIVE_HTTP_REQUEST_CLI=${COOPHIVE_HTTP_REQUEST_CLI:="curl"}

version=${version:="v0.5.4"}


detect_os_info() {
  OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && export OSARCH
  OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && export OSNAME;

  if  [ "$OSNAME" = "unsupported_os" ] || [ "$OSARCH" = "unsupported_arch" ]; then
    echo "Unsupported OS or architecture"
    echo "Checkout if our latest releases support $OSARCH_$OSNAME: https://github.com/CoopHive/hive/releases/latest"
    exit 1
  fi

}


install_hive() {
  echo "installing hive:$version"
  rurl=https://github.com/CoopHive/hive/releases/download/$version/hive-$OSNAME-$OSARCH
#  echo $rurl
  curl -sSL -o hive $rurl
  chmod +x hive
  ./hive version
#  read -p "Do you want to install Hive? (y/n): " choice
#    case "$choice" in
#      y|Y )
#        sudo cp hive /usr/local/bin/
#        ;;
#      n|N )
#        echo "Hive installation canceled."
#        ;;
#      * )
#        echo "Invalid choice. Please enter y or n."
#        ;;
#    esac
#  sudo cp hive /usr/local/bin/

}

install_bacalhau() {
#  bVersion=v1.2.1
#  wget https://github.com/bacalhau-project/bacalhau/releases/download/$bVersion/bacalhau_$bVersion_$OSNAME-$OSARCH.tar.gz
#  tar xfv bacalhau_$bVersion_$OSNAME-$OSARCH.tar.gz
#  mv bacalhau /usr/local/bin

  # Install bacalhau using curl

  curl -sL https://get.bacalhau.org/install.sh | bash
}

main() {
  detect_os_info

  if [ "$1" = "all" ]; then
    install_hive
    install_bacalhau
  elif [ "$1" = "bacalhau" ]; then
    install_bacalhau
  elif [ "$1" = "hive" ]; then
    install_hive
  else
    echo "Usage: $0 [all|bacalhau|hive]"
    exit  1
  fi
}


main "$@" || exit  1

