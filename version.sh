VERSION=$(git rev-parse HEAD)
VERSION_SHORT=$(git rev-parse --short HEAD)
TAG=$(git describe --tags)

cat <<EOF > version.go
package main

const (
 VERSION_GIT_TAG = "$TAG"
 VERSION_GIT_SHORT = "$VERSION_SHORT"
 VERSION_GIT = "$VERSION"
)
EOF
