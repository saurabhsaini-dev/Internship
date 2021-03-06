terraform {
  required_providers {
    zoom = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  token = var.jwt_token
}

resource "zoom_user" "ex" { 
  email = "thsaurabhsaini@gmail.com"
  firstname = "Saurabh"
  lastname = "Saini"
}
