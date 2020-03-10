terraform_deploy "system" {
  source {
    owner    = "terrafire"
    repo     = "terraform"
    path     = "system/"
    revision = "xxxx"
  }

  params {
    workspace   = "dev"
    vars = {
      "package_revision" = "xxx"
    }
  }
}

