#!/bin/sh

pushd /tmp

git clone https://aur.archlinux.org/linux_logo.git
pushd linux_logo

makepkg -si

popd
popd
