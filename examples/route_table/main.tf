resource "structure-deploy_route_table" "this" {
  alias        = "gmsp1zu2vntcomhubgene001-rtb97"
  name         = "gmsp1zu2vntcomhubgene001-rtb97"
  description  = "rtb-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "routetable"
  dependencies = {
    cloudprovider = "c2555fc4-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "6e0ad7e1-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "56603be7-18bd-473b-92a3-640a20bc4620"
    resourcegroup = "b71b4a40-d060-47e4-8a1f-dd1dc45e633d"
  }
  status            = "Done"
  route_propagation = "Enabled"

}

data "structure-deploy_route_table" "this" {
  name = "gmsp1zu2vntcomhubgene001-rtb97"
}
