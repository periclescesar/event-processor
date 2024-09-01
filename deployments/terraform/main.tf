resource "rabbitmq_vhost" "vhost" {
  name = "event-processor"
}

resource "rabbitmq_exchange" "main_exchange" {
  name  = "events.exchange"
  vhost = rabbitmq_vhost.vhost.name

  settings {
    type        = "direct"
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_queue" "main_queue" {
  name  = "events"
  vhost = rabbitmq_vhost.vhost.name

  settings {
    durable     = true
    auto_delete = false
    arguments = {
      "x-queue-type" : "classic",
      "x-dead-letter-exchange": rabbitmq_exchange.main_dlq_exchange.name,
      "x-dead-letter-routing-key": "",
    }
  }
}

resource "rabbitmq_binding" "main_bind" {
  vhost            = rabbitmq_vhost.vhost.name
  source           = rabbitmq_exchange.main_exchange.name
  destination      = rabbitmq_queue.main_queue.name
  destination_type = "queue"

  lifecycle {
    replace_triggered_by = [rabbitmq_queue.main_queue, rabbitmq_exchange.main_exchange]
  }
}

# DLQ
resource "rabbitmq_exchange" "main_dlq_exchange" {
  name  = "events.dlq.exchange"
  vhost = rabbitmq_vhost.vhost.name

  settings {
    type        = "direct"
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_queue" "main_dlq_queue" {
  name  = "events.dlq"
  vhost = rabbitmq_vhost.vhost.name

  settings {
    durable     = true
    auto_delete = false
    arguments = {
      "x-queue-type" : "classic",
    }
  }
}

resource "rabbitmq_binding" "main_dlq_bind" {
  vhost            = rabbitmq_vhost.vhost.name
  source           = rabbitmq_exchange.main_dlq_exchange.name
  destination      = rabbitmq_queue.main_dlq_queue.name
  destination_type = "queue"
}
