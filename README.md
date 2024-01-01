<div id="top"></div>

<!-- PROJECT SHIELDS -->
<p align="center">
<a href="https://github.com/hominsu/slink/graphs/contributors"><img src="https://img.shields.io/github/contributors/hominsu/slink.svg?style=for-the-badge" alt="Contributors"></a>
<a href="https://github.com/hominsu/slink/network/members"><img src="https://img.shields.io/github/forks/hominsu/slink.svg?style=for-the-badge" alt="Forks"></a>
<a href="https://github.com/hominsu/slink/stargazers"><img src="https://img.shields.io/github/stars/hominsu/slink.svg?style=for-the-badge" alt="Stargazers"></a>
<a href="https://github.com/hominsu/slink/issues"><img src="https://img.shields.io/github/issues/hominsu/slink.svg?style=for-the-badge" alt="Issues"></a>
<a href="https://github.com/hominsu/slink/blob/master/LICENSE"><img src="https://img.shields.io/github/license/hominsu/slink.svg?style=for-the-badge" alt="License"></a>
<a href="https://github.com/hominsu/slink/actions/workflows/docker-deploy.yml"><img src="https://img.shields.io/github/actions/workflow/status/hominsu/slink/docker-deploy.yml?branch=main&style=for-the-badge" alt="Deploy"></a>
</p>

<!-- PROJECT LOGO -->
<br/>
<div align="center">
<h3 align="center">slink</h3>

  <p align="center">
    <br/>
    <a href="https://github.com/hominsu/slink"><strong>Explore the docs »</strong></a>
    <br/>
    <br/>
    <a href="https://slink.hauhau.cn">View Demo</a>
    ·
    <a href="https://github.com/hominsu/slink/issues">Report Bug</a>
    ·
    <a href="https://github.com/hominsu/slink/issues">Request Feature</a>
  </p>
</div>

## Build & Contributing

### Backend

Before building, you need to have `GO >= 1.18`, [Buf CLI](https://docs.buf.build/installation). If you are developing on windows, use [scoop](https://github.com/ScoopInstaller/Scoop) to install `busybox` and `make`

#### Clone this repository

```bash
git clone https://github.com/hominsu/slink.git
```

#### Initial Workspace

```bash
go work init && go work use -r ./app && go mod tidy
```

#### Install dependencies

```bash
make init
```

#### Generate other code

```bash
make api && make conf && make ent && make wire
```

#### Compile

```bash
make build
```

### Frontend

```bash
cd web
pnpm install
pnpm dev
```
