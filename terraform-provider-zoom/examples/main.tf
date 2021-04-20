terraform {
  required_providers {
    zoom = {
      version = "0.1"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {}

resource "create_user" "ex" {
  user {
    email = "thsaurabhsaini@gmail.com"
    firstname = "Saurabh"
    lastname = "Saini"
  }
}
