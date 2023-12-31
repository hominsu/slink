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
    "slink-backend",
    "slink-frontend",
  ]
}

target "slink-backend" {
  context    = "."
  dockerfile = "app/slink/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    VERSION           = "$(VERSION)"
    APP_RELATIVE_PATH = "slink/service"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/slink:backend-${VERSION}" : "",
    "${REPO}/slink:backend",
  ]
  platforms = ["linux/amd64", "linux/arm64"]
}

target "slink-frontend" {
  context    = "."
  dockerfile = "web/Dockerfile"
  args = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    VERSION           = "$(VERSION)"
  }
  secret = [
    "type=env,id=NEXT_PUBLIC_REDIRECT_URL",
  ]
  tags = [
    notequal("", VERSION) ? "${REPO}/slink:frontend-${VERSION}" : "",
    "${REPO}/slink:frontend",
  ]
  platforms = ["linux/amd64", "linux/arm64"]
}