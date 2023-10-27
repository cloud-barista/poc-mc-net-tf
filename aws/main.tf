# Define the required version of Terraform and the providers that will be used in the project
terraform {
  required_version = "1.5.5"

  required_providers {
    # AWS provider is specified with its source and version
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.21"
    }
  }
}

# Provider block for AWS specifies the configuration for the provider
provider "aws" {
  region = "ap-northeast-2"
}
