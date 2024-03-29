"use client"

import { useEffect, useState } from "react";
import { useForm } from "react-hook-form"
import { useAuth } from "@/contexts/AuthContext"
import { useRouter } from "next/navigation"

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Textarea } from "@/components/ui/textarea"
import { Input } from "@/components/ui/input"
import { Toaster } from "@/components/ui/toaster"
import { useToast } from "@/components/ui/use-toast"
import { Button } from "@/components/ui/button";
import { Navbar } from "@/components/Navbar"
import { jwtDecode } from "jwt-decode"
import { JWTType } from "@/types/jwt.type"
import { UserType } from "@/types/user.type"
import axios from "axios"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { baseUrl } from "@/env";
import { MagnifyingGlassIcon } from "@radix-ui/react-icons";
import { PostsList } from "@/components/PostsList";

const formSchema = z.object({
  postTitle: z.string().min(1).max(50, {
    message: "Post title must be up to 50 characters"
  }),
  postContent: z.string().min(1).max(300, {
    message: "Post content must be at least 300 characters"
  }),
})

export default function Home() {

  const router = useRouter()
  const { token } = useAuth()
  const { toast } = useToast()

  const [user, setUser] = useState<UserType>()
  const [showToast, setShowToast] = useState<boolean>(false)

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      postTitle: "",
      postContent: ""
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    axios({
      method: "post",
      url: `${baseUrl}/posts`,
      headers: {
        Authorization: `Bearer ${token}`
      },
      data: {
        title: values.postTitle,
        content: values.postContent
      }
    }).then((response) => {
      toast({
        title: "Published post",
        description: "Your post has just been published"
      })
      console.log(response)
    })
  }

  useEffect(() => {
    if (!token) {
      return router.push("/login")
    }

    const id = jwtDecode<JWTType>(String(token)).userID
    axios.get(`${baseUrl}/users/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    }).then(({ data }) => {
      const userData: UserType = {
        id: data.id,
        email: data.email,
        name: data.name,
        username: data.username,
        createdAt: new Date(data.createdAt)
      }
      setUser(userData)
      console.log(userData)
    }).catch((error) => {
      console.error(error)
    })
  }, [])


  return (

    <>
      {
        user && (
          <div>
            <Navbar user={user} />
            <main className="w-screen flex lg:flex-row flex-col">
              <aside className="lg:w-5/12 w-full p-6">
                <div className="lg:hidden md:hidden flex relative items-center w-full mb-10">
                  <MagnifyingGlassIcon className="absolute left-2 size-6 text-muted-foreground" />
                  <Input placeholder="Search user" className="pl-8" />
                </div>
                <Form {...form}>
                  <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                      control={form.control}
                      name="postTitle"
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Title</FormLabel>
                          <FormControl>
                            <Input placeholder="Type your post title" {...field} />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name="postContent"
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Content</FormLabel>
                          <FormControl>
                            <Textarea className="resize-none" placeholder="Type your post content" {...field} />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <Button className="w-full" type="submit">Post</Button>
                  </form>
                </Form>
              </aside>
              <div className="lg:7/12">
                <PostsList user={user}/>
              </div>
            </main>
            <Toaster />
          </div>
        )
      }
    </>
  );
}
