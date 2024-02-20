#!/bin/sh
# shellcheck shell=dash

set -u

detect_os_info() {
  OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && export OSARCH
  OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && export OSNAME;

  if  [ "$OSNAME" = "unsupported_os" ] || [ "$OSARCH" = "unsupported_arch" ]; then
    echo "Unsupported OS or architecture"
    echo "Checkout if our latest releases support $OSARCH_$OSNAME: https://github.com/CoopHive/hive/releases"
    exit 1
  fi

}

install_hive() {
  version=v0.10.0
  curl -sSL -o hive https://github.com/CoopHive/hive/releases/download/$version/hive-$OSNAME-$OSARCH
  chmod +x hive
  ./hive version
#  read -p "Do you want to install Hive? (y/n): " choice
#    case "$choice" in
#      y|Y )
#        sudo mv hive /usr/local/bin/hive
#        ;;
#      n|N )
#        echo "Hive installation canceled."
#        ;;
#      * )
#        echo "Invalid choice. Please enter y or n."
#        ;;
#    esac
  sudo mv hive /usr/local/bin/hive

}

install_bacalhau() {
  bVersion=v1.2.1
  wget https://github.com/bacalhau-project/bacalhau/releases/download/$bVersion/bacalhau_$bVersion_$OSNAME-$OSARCH.tar.gz
  tar xfv bacalhau_$bVersion_$OSNAME-$OSARCH.tar.gz
  mv bacalhau /usr/local/bin

#  curl -sL https://get.bacalhau.org/install.sh | bash
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

GITHUB_ORG=CoopHive
GITHUB_REPO=hive

getLatestRelease() { //TODO:

    # /latest ignores pre-releases, see https://docs.github.com/en/rest/releases/releases#get-the-latest-release
    local tag_regex='v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)*'
    if [ "$PRE_RELEASE" == "true" ]; then
        echo "Installing most recent pre-release version..."
        local hiveReleaseUrl="https://api.github.com/repos/${GITHUB_ORG}/${GITHUB_REPO}/releases?per_page=1"
    else
        local hiveReleaseUrl="https://api.github.com/repos/${GITHUB_ORG}/${GITHUB_REPO}/releases/latest"
    fi
    local latest_release=""

    if [ "$BACALHAU_HTTP_REQUEST_CLI" == "curl" ]; then
        latest_release=$(curl -s $hiveReleaseUrl  | grep \"tag_name\" | grep -E -i "\"$tag_regex\"" | awk 'NR==1{print $2}' | sed -n 's/\"\(.*\)\",/\1/p')
    else
        latest_release=$(wget -q --header="Accept: application/json" -O - $hiveReleaseUrl | grep \"tag_name\" | grep -E -i "^$tag_regex$" | awk 'NR==1{print $2}' |  sed -n 's/\"\(.*\)\",/\1/p')
    fi

    ret_val=$latest_release
}