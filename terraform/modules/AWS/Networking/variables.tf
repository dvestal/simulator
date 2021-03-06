variable "vpc_cidr" {
  description = "cidr block for vpc"
  default     = "172.31.0.0/16"
}

variable "public_subnet_cidr" {
  description = "cidr range for public subnet"
  default     = "172.31.1.0/24"
}

variable "private_subnet_cidr" {
  description = "cidr range for private subnet"
  default     = "172.31.2.0/24"
}

variable "default_tags" {
  description = "Default tags for all resources"
  type        = "map"
}

