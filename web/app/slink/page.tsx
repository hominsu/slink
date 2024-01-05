import { LinkShorter } from './link-shorter'

export default async function SlinkPage() {
  return (
    <div className="h-full flex-1 flex-col space-y-8 p-2 md:flex">
      <h2 className="pl-6 text-2xl font-bold tracking-tight">Welcome back!</h2>
      <div className="flex items-center justify-center">
        <div className="w-full max-w-sm p-4">
          <LinkShorter className="rounded-lg border shadow-lg" />
        </div>
      </div>
    </div>
  )
}
