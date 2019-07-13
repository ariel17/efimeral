#!/bin/sh

pushd /tmp

git clone https://aur.archlinux.org/shellinabox-git.git
pushd shellinabox-git

makepkg -si

popd
popd
