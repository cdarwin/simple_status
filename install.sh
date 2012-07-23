#!/bin/sh -e

##
# This script makes certain assumptions about your deployment of SimpleStatus
# 1) You are deploying to at least Ubuntu 12.04 LTS
# 2) the user running the script has nopasswd access to sudo
# 3) using upstart to manage the daemonizing
# This script should be viewed as a template for your own deployment and should
# not be run blindly

##
# Set up build directory
chdir $HOME
mkdir ssd
git clone https://github.com/cdarwin/simple_status.git ssd
cd ssd
export DIR=$(pwd)

## 
# We know where generate_cert.go lives when installed from ppa
# but, if you want to find it yourself, uncomment the line below
# and use the $CRT variable in the build stage
#export CRT=$(sudo find / -name generate_cert.go)

##
# Install Go if it isn't already
if [ ! $(go version) ]; then
  sudo /bin/sh -c 'echo "y" | add-apt-repository ppa:gophers/go'
  sudo apt-get -yf update
  sudo apt-get -yf install golang-stable
fi

##
# Build package, generate certs, and modify permissions
go build simple_statusd.go
#go build $CRT
go build /usr/lib/go/src/pkg/crypto/tls/generate_cert.go
./generate_cert
sudo chown :www-data key.pem
sudo chmod g+r key.pem

##
# This is just one way to deal changing the default configuration of the package
# we're using upstart here and setting some runtime flags
sed -i -e 's,TLS=,&"-ssl",' -e 's,PORT=,&"-p :9090",' -e 's,TOKEN=,&"-t foobarbaz",' \
  -e "s,\(DIR=\).*,\1\"$DIR\"," simple_statusd.conf
sudo cp simple_statusd.conf /etc/init/simple_statusd.conf
sudo start simple_statusd
