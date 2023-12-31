import { ConnectError, createPromiseClient } from '@connectrpc/connect'
import { createConnectTransport } from '@connectrpc/connect-web'

import { ShortLinkService as ConnectShortLinkService } from '@/lib/api/slink/service/v1/slink_connect'

export interface IShortLinkService {
  createShortLink(link: string): Promise<string>
}

export class ShortLinkService implements IShortLinkService {
  private client

  constructor() {
    this.client = createPromiseClient(
      ConnectShortLinkService,
      createConnectTransport({
        baseUrl: '/api',
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

export const shortLinkService = new ShortLinkService()
