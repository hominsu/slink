export const siteConfig = {
  name: 'slink',
  url: 'https://slink.hauhau.cn',
  ogImage: 'https://slink.hauhau.cn/og.png',
  description:
    'Compact, user-friendly short link service. Open Source. Transform long URLs into short, manageable links easily. Ideal for sharing and tracking.',
  links: {
    author: 'https://homing.so/about',
    twitter: 'https://twitter.com/is_homingso',
    github: 'https://github.com/hominsu/slink',
  },
  site: {
    redirectURL: process.env.NEXT_PUBLIC_REDIRECT_URL,
    grpcEndpoint: process.env.NEXT_GRPC_ENDPOINT,
  },
}

export type SiteConfig = typeof siteConfig
