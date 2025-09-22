#########################
#   Main app variable   #
#########################
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


