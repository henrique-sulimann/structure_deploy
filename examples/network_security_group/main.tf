resource "structure-deploy_network_security_group" "this" {
  alias        = "gmsp1zu2nsgcomhubgene099"
  name         = "gmsp1zu2nsgcomhubgene099"
  description  = "nsg-zu2-hub-test-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "networksecuritygroup"
  dependencies = {
    cloudprovider = "111111-8e21-4ad1-bcfa-111111"
    tenant        = "111111-734d-48af-ae39-111111"
    subscription  = "111111-18bd-473b-92a3-111111"
    resourcegroup = "111111-d060-47e4-8a1f-1111111"
  }
  status = "Done"

}

data "structure-deploy_network_security_group" "this" {
  name = "gmsp1zu2nsgcomhubgene099"
}
