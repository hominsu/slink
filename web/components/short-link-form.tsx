import React, { FC, useState } from 'react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

interface ShortLinkFormProps {
  onSubmit: (link: string, resetInput: () => void) => void
}

export const ShortLinkForm: FC<ShortLinkFormProps> = ({ onSubmit }) => {
  const [inputLink, setInputLink] = useState('')

  const resetInput = () => {
    setInputLink('')
  }

  const handleSubmit = () => {
    onSubmit(inputLink, resetInput)
  }

  return (
    <div className="flex space-x-2">
      <Input
        placeholder="https://www.hauhau.cn"
        value={inputLink}
        onChange={(e) => setInputLink(e.target.value)}
      />
      <Button className="shrink-0" onClick={handleSubmit}>
        Submit
      </Button>
    </div>
  )
}
