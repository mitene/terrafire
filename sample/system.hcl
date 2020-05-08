workspace "system" {
  source "github" {
    owner = "terrafire"
    repo  = "terraform"
    path  = "system/"
    ref   = "xxxx"
  }

  workspace   = "dev"
  vars = {
    "package_revision" = "xxx"
  }
}

