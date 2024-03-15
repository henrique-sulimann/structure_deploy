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
  name = "oracle cloud"
}
