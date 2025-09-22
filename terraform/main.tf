terraform {
  required_providers {
    docker = { source = "kreuzwerker/docker", version = "~> 3.0" }
  }
}


provider "docker" {}

resource "docker_network" "app" {
  name = "app_net"
}


resource "docker_image" "app" {
  name         = var.image
  keep_locally = true
}

resource "docker_container" "app" {
  name  = var.container_name
  image = docker_image.app.image_id
  restart = "unless-stopped"

networks_advanced { name = docker_network.app.name }

  ports {
    internal = var.container_port
    external = var.host_port
    protocol = "tcp"
  }

}