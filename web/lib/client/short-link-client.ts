import { ConnectError, createPromiseClient } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'

import { siteConfig } from '@/config/site'
import { ShortLinkService } from '@/lib/api/slink/service/v1/slink_connect'

export interface IShortLinkClient {
  createShortLink(link: string): Promise<string>
}

export class ShortLinkClient implements IShortLinkClient {
  private client

  constructor() {
    this.client = createPromiseClient(
      ShortLinkService,
      createGrpcTransport({
        baseUrl: siteConfig.site.grpcEndpoint as string,
        httpVersion: '2',
        interceptors: [],
      })
    )
  }

  async createShortLink(link: string): Promise<string> {
    try {
      const response = await this.client.createShortLink({ link })
      return response.key
    } catch (err) {
      let errorMessage = 'Error creating short link'
      if (err instanceof ConnectError) {
        errorMessage += `: ${err.code} - ${err.message}`
      } else if (err instanceof Error) {
        errorMessage += `: ${err.message}`
      }
      throw new Error(errorMessage)
    }
  }
}

export const shortLinkClient = new ShortLinkClient()
