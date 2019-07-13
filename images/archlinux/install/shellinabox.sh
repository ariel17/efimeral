#!/bin/sh

pushd /tmp/shellinabox-git

pacman -U shellinabox-git-*.pkg.tar* --noconfirm

popd
