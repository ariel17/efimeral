#!/bin/sh

pushd /tmp/linux_logo

pacman -U linux_logo*.pkg.tar* --noconfirm

popd
