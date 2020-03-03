terrafire {
  backend "s3" {
    bucket = "state_bucket"
    key    = "state_file"
  }
}