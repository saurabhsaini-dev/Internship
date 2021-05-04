terraform {
  required_providers{
    zoom ={
      version ="0.2"
      source = "hashicorp.com/edu/zoom"
    }
  }
}
provider "zoom" {
  auth_token= "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk"
}
/*
data "zoom_users" "all"{
  email= "ashishdhodria17@cse.iiitp.ac.in"
}

output "users" {
  value = data.zoom_users.all
}
*/
/*
resource "zoom_User_instance" "user" {
  first_name = "ashish"
  last_name  = "dhodria"
  email      = "ashishdhodria1999@gmail.com"
  type       = 1
}

output "user_instance" {
  value = zoom_User_instance.user
}
*/