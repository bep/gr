#!/usr/bin/env bash

for D in `find . -type d -not -path '*/\.*'`
do
    pushd $D
	gopherjs -m build
	popd
done

pushd interop
browserify -r ./interop-ext-module.js:ElapserExtModule > interop-ext-module-bundle.js 
browserify -x ./interop-ext-module.js interop-ext-reverse.js > interop-ext-reverse-bundle.js 

popd
