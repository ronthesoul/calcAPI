terraform {
  required_providers {
    docker = { source = "kreuzwerker/docker", version = "~> 3.0" }
  }
}

variable "image" {
  type    = string
  default = "m4gapower/calcapi:1.18.0"
}

variable "container_name" {
  type    = string
  default = "calcapi"
}

variable "container_port" {
  type    = number
  default = 8080
}

variable "host_port" {
  type    = number
  default = 8080
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