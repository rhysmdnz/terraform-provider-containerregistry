package main

import (
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/authn"

)


func main() {
	tag, err := name.NewTag("australia-southeast1-docker.pkg.dev/memesnz/test/pinger")
	if err != nil {
		panic(err)
	}
	img, err := tarball.ImageFromPath("pinger-gcp.tar", nil)
	if err != nil {
		panic(err)
	}
	err = remote.Write(tag, img, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		panic(err)
	}
}
