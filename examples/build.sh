#!/usr/bin/env bash

for D in `find . -type d -not -path '*/\.*'`
do
    pushd $D
	gopherjs -m build
	popd
done


