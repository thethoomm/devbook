"use client"

import { useAuth } from '@/contexts/AuthContext'
import { useRouter } from 'next/navigation'

import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import Image from 'next/image'
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { Separator } from "@/components/ui/separator"
import { ExclamationTriangleIcon, GitHubLogoIcon } from '@radix-ui/react-icons'
import axios from 'axios'
import Link from 'next/link'
import { useState } from 'react'
import { baseUrl } from '@/env'

const formSchema = z.object({
  email: z.string().email({
    message: "Invalid email."
  }),
  password: z.string().min(3, {
    message: "Password must be at least 3 characters."
  })
})

export default function Login() {

  const router = useRouter()
  const { setToken, token } = useAuth()

  const [showAlert, setShowAlert] = useState<boolean>(false)

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: ""
    }
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    axios.post(`${baseUrl}/login`, {
      email: values.email,
      password: values.password
    }).then((response) => {
      setToken(String(response.data))
      localStorage.setItem('token', response.data)
      router.push("/")
    }).catch(function (error) {
      setShowAlert(true)
      console.log(error);
    })
  }

  return (
    <main className='w-screen h-screen flex flex-row'>
      <div className='w-full sm:flex sm:w-7/12 bg-zinc-50 items-center justify-center sm:p-12 p-6'>
        <div className='flex-1 p-3'>
          <h1 className='text-2xl text-primary font-bold text-center'>Login</h1>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-6'>
              <FormField
                control={form.control}
                name='email'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input placeholder='example@gmail.com' {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='password'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Password</FormLabel>
                    <FormControl>
                      <Input type='password' {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <p className='text-xs text-right space-x-1'>
                <span>Don&apos;t have an account?</span>
                <Link
                  className='text-primary font-medium'
                  href="/signup"
                >
                  Sign up
                </Link>
              </p>
              {
                showAlert && (
                  <Alert variant={'destructive'}>
                    <ExclamationTriangleIcon className="size-6" />
                    <AlertTitle className='ml-2'>Error</AlertTitle>
                    <AlertDescription className='ml-2'>
                      An internal server error has occurred
                    </AlertDescription>
                  </Alert>

                )
              }
              <Button className='w-full'>Login</Button>
              <Separator />
              <div className='flex-1'>
                <Button variant={"outline"} className='space-x-2 w-full'>
                  <GitHubLogoIcon className='size-4' />
                  <p>Login with Github</p>
                </Button>
              </div>
            </form>
          </Form>
        </div>
      </div>
      <div className='hidden sm:w-5/12 sm:flex'>
        <Image
          src="https://images.unsplash.com/photo-1522098543979-ffc7f79a56c4?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
          width={1000}
          height={1000}
          alt="Image"
          className="rounded-md object-cover" />
      </div>
    </main>
  )
}