terraform_deploy "system" {
  source {
    owner    = "terrafire"
    repo     = "terraform"
    path     = "system/"
    revision = "xxxx"
  }

  workspace   = "dev"

  vars = {
    "package_revision"          = "xxx"
  }
}

