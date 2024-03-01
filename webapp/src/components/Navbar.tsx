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
import { UserType } from "@/types/user.type"

interface Props {
  user: UserType
}

export function Navbar({ user }: Props) {

  return (
    <nav className="w-full bg-zinc-950 text-secondary items-center flex flex-row justify-between py-2 px-6">
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
            <Button variant={"ghost"} className="flex flex-row items-center space-x-2">
              <Avatar className="size-8">
                <AvatarImage src="https://github.com/thethoomm.png" />
                <AvatarFallback>
                  {
                    user.username.slice(0,2)
                  }
                </AvatarFallback>
              </Avatar>
              <h1>{ user.username }</h1>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuItem>Preferences</DropdownMenuItem>
            <DropdownMenuItem>Support</DropdownMenuItem>
            <DropdownMenuItem>
              <p className="text-red-400">Logout</p>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </nav>
  )
}