resource "structure-deploy_virtual_network" "this" {
  alias        = "gmsp1zu2vntcomhubgene999"
  name         = "gmsp1zu2vntcomhubgene999"
  description  = "vnt-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "virtualNetwork"
  dependencies = {
    cloudprovider = "c2555fc4-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "6e0ad7e1-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "56603be7-18bd-473b-92a3-640a20bc4620"
    resourcegroup = "b71b4a40-d060-47e4-8a1f-dd1dc45e633d"
  }
  status = "done"
  cidr = [
    "107.104.32.0/22",
    "107.104.40.0/22",
    "180.41.0.0/22"
  ]

}

data "structure-deploy_virtual_network" "this" {
  name = "gmsp1zu2vntcomhubgene999"
}

