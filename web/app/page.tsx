import Link from 'next/link'

import { siteConfig } from '@/config/site'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import { Announcement } from '@/components/announcement'
import { Icons } from '@/components/icons'
import {
  PageActions,
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from '@/components/page-header'

import { LinkShorter } from './link-shorter'

export default async function IndexPage() {
  return (
    <div className="container relative">
      <PageHeader>
        <Announcement />
        <PageHeaderHeading>Shorter your link</PageHeaderHeading>
        <PageHeaderDescription>
          Compact, user-friendly short link service. Open Source. Transform long URLs into short,
          manageable links easily. Ideal for sharing and tracking.
        </PageHeaderDescription>
        <PageActions>
          <Link
            target="_blank"
            rel="noreferrer"
            href={siteConfig.links.github}
            className={cn(buttonVariants({ variant: 'outline' }))}
          >
            <Icons.gitHub className="mr-2 h-4 w-4" />
            GitHub
          </Link>
        </PageActions>
      </PageHeader>
      <div className="flex items-center justify-center">
        <div className="w-full max-w-md p-4">
          <LinkShorter className="rounded-lg border shadow-lg" />
        </div>
      </div>
    </div>
  )
}
