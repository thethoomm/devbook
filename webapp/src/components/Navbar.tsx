import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"


import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { MagnifyingGlassIcon } from "@radix-ui/react-icons"


export function Navbar() {
  return (
    <nav className="w-full bg-zinc-900 text-secondary items-center flex flex-row justify-between py-2 px-6">
      <div className="w-fit">
        <p className="font-mono font-bold">devbook.io</p>
      </div>
      <div className="flex relative items-center w-6/12">
        <MagnifyingGlassIcon className="absolute left-2 size-6 text-muted-foreground" />
        <Input placeholder="Search user" className="pl-8 border-0 bg-zinc-800" />
      </div>
      <div className="">
        <DropdownMenu>
          <DropdownMenuTrigger>
            <div>
              <Avatar className="size-8">
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
              <h1>Text</h1>
            </div>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuItem>Billing</DropdownMenuItem>
            <DropdownMenuItem>Team</DropdownMenuItem>
            <DropdownMenuItem>Subscription</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </nav>
  )
}