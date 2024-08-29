terraform {
  required_version = "1.8.2"

  required_providers {
    rabbitmq = {
      source = "cyrilgdn/rabbitmq"
      version = "1.8.0"
    }
  }
}

provider "rabbitmq" {
  endpoint = var.RABBITMQ_MANAGEMENT_ENTRYPOINT
  username = var.RABBITMQ_DEFAULT_USER
  password = var.RABBITMQ_DEFAULT_PASS
}