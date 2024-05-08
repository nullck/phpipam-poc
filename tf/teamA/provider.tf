terraform {
  required_providers {
    phpipam = {
      source = "lord-kyron/phpipam"
      version = "1.5.1"
    }
  }
}

provider "phpipam" {
  app_id   = "apiclient"
  endpoint = "http://localhost/api"
  username = "admin"
  password = "secret123"
  insecure = true
}
