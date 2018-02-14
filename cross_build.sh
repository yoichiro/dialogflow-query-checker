#!/bin/sh

if [ ! -d "dist" ]
then
    mkdir "dist"
fi
VERSION=`cat version.go | awk 'match($0, /[0-9.]+/) { print substr($0, RSTART, RLENGTH) }'`
DIST="dist/dialogflow-query-checker-${VERSION}"
rm -rf "${DIST}"
rm -f "${DIST}.zip"
gox \
    -os="darwin linux windows" \
    -arch="386 amd64" \
    -output "${DIST}/{{.OS}}_{{.Arch}}/{{.Dir}}"
zip -r "${DIST}.zip" "${DIST}"
