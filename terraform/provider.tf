provider "azurerm" {
  features {}
}
terraform {
  required_providers {
    structure-deploy = {
      version = "0.0.4-223"
      source  = "hsulimann.com/dev/structure-deploy"
    }
  }
}
provider "structure-deploy" {}
