import { Timestamp } from '@bufbuild/protobuf'
import { ConnectError, createPromiseClient } from '@connectrpc/connect'
import { createConnectTransport } from '@connectrpc/connect-web'

import { ShortLinkService as ConnectShortLinkService } from '@/lib/api/slink/service/v1/slink_connect'
import { CreateShortLinkReply } from '@/lib/api/slink/service/v1/slink_pb'
import { InputForm } from '@/lib/scheme/input-form'

export interface IShortLinkService {
  createShortLink(data: InputForm): Promise<CreateShortLinkReply>
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

  async createShortLink(data: InputForm): Promise<CreateShortLinkReply> {
    try {
      return await this.client.createShortLink({
        link: data.link,
        expireAt: Timestamp.fromDate(data.expireAt),
      })
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
