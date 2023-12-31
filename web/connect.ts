import { ConnectRouter } from '@connectrpc/connect'

import { ShortLinkService } from '@/lib/api/slink/service/v1/slink_connect'
import type { CreateShortLinkRequest } from '@/lib/api/slink/service/v1/slink_pb'
import { shortLinkClient } from '@/lib/client/short-link-client'

// eslint-disable-next-line import/no-anonymous-default-export
export default function (router: ConnectRouter) {
  router.rpc(
    ShortLinkService,
    ShortLinkService.methods.createShortLink,
    async (req: CreateShortLinkRequest) => {
      try {
        const key = await shortLinkClient.createShortLink(req.link)
        return { key }
      } catch (error) {
        console.error(error)
        throw error
      }
    }
  )
}
