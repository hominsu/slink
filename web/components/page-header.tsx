import Balance from 'react-wrap-balancer'

import { cn } from '@/lib/utils'

function PageHeader({ className, children, ...props }: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <section
      className={cn(
        'mx-auto flex max-w-[980px] flex-col items-center gap-2 py-2 md:py-4 md:pb-2 lg:py-8 lg:pb-4',
        className
      )}
      {...props}
    >
      {children}
    </section>
  )
}

function PageHeaderHeading({ className, ...props }: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    // eslint-disable-next-line jsx-a11y/heading-has-content
    <h1
      className={cn(
        'text-center text-3xl font-bold leading-tight tracking-tighter md:text-6xl lg:leading-[1.1]',
        className
      )}
      {...props}
    />
  )
}

function PageHeaderDescription({
  className,
  ...props
}: React.HTMLAttributes<HTMLParagraphElement>) {
  return (
    <Balance
      className={cn(
        'max-w-[750px] text-center text-lg text-muted-foreground sm:text-xl',
        className
      )}
      {...props}
    />
  )
}

function PageActions({ className, ...props }: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn('flex w-full items-center justify-center space-x-4 py-2 md:pb-4', className)}
      {...props}
    />
  )
}

export { PageHeader, PageHeaderHeading, PageHeaderDescription, PageActions }
