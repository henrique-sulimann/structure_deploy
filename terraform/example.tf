terraform {
  required_providers {
    ccoe-naming = {
      version = "0.0.7-2"
      source  = "hsulimann.com/dev/ccoe-naming"
    }
  }
}
provider "ccoe-naming" {}
# data "ccoe-naming_resources" "test" {}
variable "region" {
  default = [
    "eastus2",
    "eastus2"
  ]
}
resource "ccoe-naming_resources" "test" {
  count = 2
  # name        = "gmsp1zu2asbplatfogene13${count.index + 1}"
  product     = "Azure Service Bus"
  function    = "generic"
  application = "Platform"
  region      = var.region[count.index]
  env         = "production"
}
resource "ccoe-naming_vms" "test" {
  region  = "eastus2"
  env     = "production"
  product = "genericvm"
  os      = "windows"
}
output "test-resource-ccoe-naming" {
  value = ccoe-naming_resources.test
}
output "test-vm-ccoe-naming" {
  value = ccoe-naming_vms.test
}

