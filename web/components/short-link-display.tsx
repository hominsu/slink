import React, { FC } from 'react'
import { InfoCircledIcon } from '@radix-ui/react-icons'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip'

interface ShortLinkDisplayProps {
  link: string
}

export const ShortLinkDisplay: FC<ShortLinkDisplayProps> = ({ link }) => {
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

  return (
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
  )
}
