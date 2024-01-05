import * as z from 'zod'

export const InputFormSchema = z.object({
  link: z.string().url({
    message: 'Please enter a valid URL.',
  }),
  expireAt: z.date({
    required_error: 'An expiration date is required',
  }),
})

export type InputForm = z.infer<typeof InputFormSchema>
