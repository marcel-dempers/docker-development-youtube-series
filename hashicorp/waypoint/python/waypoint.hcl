project = "example-nodejs"

app "example-nodejs" {
  labels = {
      "service" = "example-nodejs",
      "env" = "dev"
  }

  build {
    use "pack" {}
    registry {
        use "docker" {
          image = "aimvector/waypoint-python-example"
          tag = "1"
          local = false
        }
    }
 }

  deploy { 
    use "kubernetes" {
    probe_path = "/"
    }
  }

  release {
    use "kubernetes" {
      port = 5000
    }
  }
}