# Region to deploy into
variable "aws_region" {
  type    = string
  default = "us-west-2"
}

# ECR & ECS settings
variable "ecr_repository_name" {
  type    = string
  default = "ecr_service"
}

variable "service_name" {
  type    = string
  default = "RideBookingService"
}

variable "container_port" {
  type    = number
  default = 8080
}

variable "ecs_count" {
  type    = number
  default = 10
}

# How long to keep logs
variable "log_retention_days" {
  type    = number
  default = 7
}

# Auto Scaling configuration
variable "ecs_min_capacity" {
  description = "Minimum number of ECS tasks to keep running"
  type        = number
  default     = 1
}

variable "ecs_max_capacity" {
  description = "Maximum number of ECS tasks to scale out to"
  type        = number
  default     = 4
}
variable "cpu_target" {
  description = "Target average CPU utilization for ECS auto scaling"
  type        = number
  default     = 70
}