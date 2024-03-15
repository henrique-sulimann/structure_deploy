resource "structure-deploy_subscription" "this" {
  alias        = "gcld1glbsubplatfogene021"
  name         = "gcld1glbsubplatfogene021"
  description  = "cl - gcl - dev 2"
  version      = "v1alpha1"
  release_date = "2023-06-16"
  labels = {
    cloud  = "az"
    teste  = "123"
    teste1 = "1234"
  }
  kind = "subscription"
  dependencies = {
    cloudprovider = "c2555fc4-8e21-4ad1-bcfa-bebc4abf56db"
    tenant        = "6e0ad7e1-734d-48af-ae39-f1027fa7ca7f"
  }
  subscription_id = "b9121c3b-8bee-49e9-89ec-97df332a5130"
}
data "structure-deploy_subscription" "this" {
  id = "93b2c4c3-2973-4b6f-a15e-4f569882c6e6"
}

