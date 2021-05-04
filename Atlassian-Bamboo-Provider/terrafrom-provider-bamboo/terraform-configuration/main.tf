terraform {
  required_providers{
    bamboo = {
      version ="0.2"
      source = "hashicorp.com/edu/bamboo"
    }
  }
}
provider "zoom" {
  auth_token= "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk"
}
 
resource "bamboo_user_instance" "user" {
   
}

