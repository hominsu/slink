'use client'

import React, { FC, useState } from 'react'

import { siteConfig } from '@/config/site'
import { shortLinkService } from '@/lib/service/short-link-service'
import { CreateFullUrl } from '@/lib/utils'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import { ShortLinkDisplay } from '@/components/short-link-display'
import { ShortLinkForm } from '@/components/short-link-form'

interface LinkShorterProps extends React.HTMLAttributes<HTMLDivElement> {}

export const LinkShorter: FC<LinkShorterProps> = ({ className, ...props }: LinkShorterProps) => {
  const [shortenedLink, setShortenedLink] = useState<string | null>(null)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)

  const handleSubmit = async (inputLink: string, resetInput: () => void) => {
    try {
      const key = await shortLinkService.createShortLink(inputLink)
      const fullUrl = CreateFullUrl(key, siteConfig.site.redirectURL as string)
      if (fullUrl) {
        resetInput()
        setShortenedLink(fullUrl)
        setErrorMessage(null)
      }
    } catch (error) {
      setErrorMessage(error instanceof Error ? error.message : 'An unknown error occurred')
    }
  }

  return (
    <Card className={className} {...props}>
      <CardHeader>
        <CardTitle>Create Your Short Link</CardTitle>
        <CardDescription>Enter a URL to generate a short and shareable link.</CardDescription>
      </CardHeader>
      <CardContent>
        <ShortLinkForm onSubmit={handleSubmit} />
        {errorMessage && (
          <>
            <Separator className="my-4" />
            <div className="my-4 text-sm text-red-500">{errorMessage}</div>
          </>
        )}
        {shortenedLink && !errorMessage && (
          <>
            <Separator className="my-4" />
            <div className="my-4 text-sm text-zinc-500 dark:text-zinc-400">
              Your short link is ready, click the link below to copy 👏 !
            </div>
            <ShortLinkDisplay link={shortenedLink} />
          </>
        )}
      </CardContent>
    </Card>
  )
}
