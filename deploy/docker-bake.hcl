variable "REPO" {
  default = "hominsu"
}

variable "AUTHOR_NAME" {
  default = "hominsu"
}

variable "AUTHOR_EMAIL" {
  default = "hominsu@foxmail.com"
}

variable "VERSION" {
  default = ""
}

group "default" {
  targets = [
    "slink-service",
  ]
}

target "slink-service" {
  context    = "."
  dockerfile = "app/slink/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    VERSION           = "$(VERSION)"
    APP_RELATIVE_PATH = "slink/service"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/slink:${VERSION}" : "",
    "${REPO}/slink:latest",
  ]
  platforms = ["linux/amd64", "linux/arm64"]
}