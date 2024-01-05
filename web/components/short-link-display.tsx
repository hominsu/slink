import React, { FC } from 'react'
import { InfoCircledIcon } from '@radix-ui/react-icons'
import { format, formatDistanceToNow } from 'date-fns'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip'

interface ShortLinkDisplayProps {
  link: string
  expireAt: Date
}

export const ShortLinkDisplay: FC<ShortLinkDisplayProps> = ({ link, expireAt }) => {
  const copyToClipboard = (link: string) => {
    navigator.clipboard.writeText(link).then(
      () =>
        toast('Link Copied', {
          description: 'Link has been copied to your clipboard.',
          action: {
            label: 'Got it',
            onClick: () => {},
          },
        }),
      () =>
        toast('Copy Failed', {
          description: 'Unable to copy the link to your clipboard.',
          action: {
            label: 'Close',
            onClick: () => {},
          },
        })
    )
  }

  const formattedExpireAt = formatDistanceToNow(expireAt, { addSuffix: true })
  const fullExpireAt = format(expireAt, 'PPpp')

  return (
    <>
      <div className="my-4 text-sm text-zinc-500 dark:text-zinc-400">
        Your short link is ready, click the link below to copy 👏 !
      </div>
      <div className="flex items-center justify-center">
        <Tooltip>
          <TooltipTrigger>
            <Badge variant="secondary" onClick={() => copyToClipboard(link)}>
              {link}
              <InfoCircledIcon className="ml-2 h-4 w-4" />
            </Badge>
          </TooltipTrigger>
          <TooltipContent>Click to copy</TooltipContent>
        </Tooltip>
      </div>
      <div className="my-4 text-sm text-zinc-500 dark:text-zinc-400">
        <Tooltip>
          This link will expire
          <TooltipTrigger>
            <span className="ml-1 underline">{formattedExpireAt}.</span>
          </TooltipTrigger>
          <TooltipContent>{fullExpireAt}</TooltipContent>
        </Tooltip>
      </div>
    </>
  )
}
