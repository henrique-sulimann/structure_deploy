resource "structure-deploy_resource_group" "this" {
  alias        = "gmsp1zu2rsgcomhubcrit096"
  name         = "gmsp1zu2rsgcomhubcrit096"
  description  = "rsg-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "resourceGroup"
  dependencies = {
    cloudprovider = "c2555fc4-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "6e0ad7e1-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "56603be7-18bd-473b-92a3-640a20bc4620"
  }
  continent = "HUB Americas"
}

data "structure-deploy_resource_group" "this" {
  id = "0e2d6375-c2d5-451c-a4d7-82107fb3c931"
}
