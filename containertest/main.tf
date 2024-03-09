terraform {
  required_providers {
    containerregistry = {
      version = "0.2.0"
      source  = "registry.terraform.io/rhysmdnz/containerregistry"
    }
  }
}


provider "containerregistry" {}

resource "containerregistry_resource" "mine" {
  image_tarball      = "containertest/pinger-gcp.tar"
  image_tarball_hash = filebase64sha256("containertest/pinger-gcp.tar")
  remote_tag         = "australia-southeast1-docker.pkg.dev/memesnz/test/pinger:main"
}
