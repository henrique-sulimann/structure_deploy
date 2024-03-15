resource "structure-deploy_subscription" "this" {
  alias       = "hsxd1zu2subplatfogene002"
  name        = "hsxd1zu2subplatfogene002"
  description = "cl - hsx - dev 2"
  version     = "v1alpha1"
  labels = {
    cloud = "az"
  }
  subscription_id = "11111111-8bee-49e9-89ec-123456789"
  kind            = "subscription"
  dependencies = {
    cloudprovider = "1111111-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "1111111-734d-48af-ae39-f1027fa7ca7f"
  }
  properties = {
    settings = {
      "subscription_id" = "11111111-8bee-49e9-89ec-123456789"
    }
  }
}

resource "structure-deploy_subscription" "this" {
  alias        = "hsxd1zu2subplatfogene021"
  name         = "hsxd1zu2subplatfogene021"
  description  = "cl - hsx - dev 2"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    cloud  = "az"
    teste  = "123"
    teste1 = "1234"
  }
  kind = "subscription"
  dependencies = {
    cloudprovider = "1111111-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "1111111-734d-48af-ae39-f1027fa7ca7f"
  }
  subscription_id = "1111111-8bee-49e9-89ec-111111111"
}
data "structure-deploy_subscription" "this" {
  name = "hsxd1zu2subplatfogene002"
}
output "test_data_source" {
  value = structure-deploy_subscription.this
}
resource "structure-deploy_resource_group" "this" {
  alias        = "hxsp1zu2rsgcomhubcrit096"
  name         = "hxsp1zu2rsgcomhubcrit096"
  description  = "rsg-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "resourceGroup"
  dependencies = {
    cloudprovider = "microsoft azure"
    tenant        = "henriquesulimann.onmicrosoft.com"
    subscription  = "hxsp1zu2subcomhubgene001"
  }
  continent = "HUB Americas"
}
resource "structure-deploy_network_security_group" "this" {
  alias        = "hxsp1zu2nsgcomhubgene099"
  name         = "hxsp1zu2nsgcomhubgene099"
  description  = "nsg-zu2-hub-test-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "networksecuritygroup"
  dependencies = {
    cloudprovider = "1111111-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "1111111-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "1111111-18bd-473b-92a3-640a20bc4620"
    resourcegroup = "1111111-d060-47e4-8a1f-dd1dc45e633d"
  }
  status = "Done"

}
resource "structure-deploy_route_table" "this" {
  alias        = "hxsp1zu2vntcomhubgene001-rtb97"
  name         = "hxsp1zu2vntcomhubgene001-rtb97"
  description  = "rtb-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "routetable"
  dependencies = {
    cloudprovider = "1111111-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "1111111-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "1111111-18bd-473b-92a3-640a20bc4620"
    resourcegroup = "1111111-d060-47e4-8a1f-dd1dc45e633d"
  }
  status            = "Done"
  route_propagation = "Enabled"

}
resource "structure-deploy_virtual_network" "this" {
  alias        = "hxsp1zu2vntcomhubgene999"
  name         = "hxsp1zu2vntcomhubgene999"
  description  = "vnt-zu2-hub-test-terraform-provider"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud"     = "az",
    "continent" = "hub_americas"
  }
  kind = "virtualNetwork"
  dependencies = {
    cloudprovider = "1111111-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "1111111-734d-48af-ae39-f1027fa7ca7f"
    subscription  = "1111111-18bd-473b-92a3-640a20bc4620"
    resourcegroup = "1111111-d060-47e4-8a1f-dd1dc45e633d"
  }
  status = "done"
  cidr = [
    "107.104.32.0/22",
    "107.104.40.0/22",
    "180.41.0.0/22"
  ]

}
resource "structure-deploy_cloud_provider" "this" {
  alias        = "oci"
  name         = "oracle cloud"
  description  = "oracle cloud"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    "cloud" : "oci"
  }
  kind = "cloudprovider"
}
data "structure-deploy_cloud_provider" "this" {
  name = "amazon web services"
}
output "test_cloud_provider_data_source" {
  value = data.structure-deploy_cloud_provider.this
}
data "structure-deploy_network_security_group" "this" {
  name = "hxsp1zu2nsgcomhubgene099"
}
output "test_network_security_group_data_source" {
  value = data.structure-deploy_network_security_group.this
}
data "structure-deploy_route_table" "this" {
  name = "hxsp1zu2vntcomhubgene001-rtb25"
}
output "test_route_table_data_source" {
  value = data.structure-deploy_route_table.this
}
data "structure-deploy_virtual_network" "this" {
  name = "hxsp1zu2vntcomhubgene999"
}
output "test_virtual_network_data_source" {
  value = data.structure-deploy_virtual_network.this
}
data "structure-deploy_resource_group" "this" {
  name = "11111-c2d5-451c-a4d7-111111"
}
output "test_resource_group_data_source" {
  value = data.structure-deploy_resource_group.this
}
