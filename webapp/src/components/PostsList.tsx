import { useEffect, useState } from "react";
import { ScrollArea } from "@/components/ui/scroll-area";
import { useAuth } from "@/contexts/AuthContext";
import { baseUrl } from "@/env";
import { PostType } from "@/types/post.type";
import axios from "axios";
import { UserType } from "@/types/user.type";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "./ui/button";
import { ArrowUpIcon } from "@radix-ui/react-icons";

interface Props {
  user: UserType
}


export function PostsList({ user }: Props) {

  const { token } = useAuth()

  const [posts, setPosts] = useState<PostType[]>()
  const [isLiked, setIsLiked] = useState<number>()

  function handleLike(postId: number) {
    setIsLiked(!isLiked)

    if (isLiked) {
      axios({
        url: `${baseUrl}/posts/${postId}/like`,
        method: "post",
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
    } else {
      axios({
        url: `${baseUrl}/posts/${postId}/dislike`,
        method: "post",
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
    }
  }

  useEffect(() => {
    axios({
      url: `${baseUrl}/posts`,
      method: "get",
      headers: {
        Authorization: `Bearer ${token}`
      }
    }).then(({ data }) => {
      setPosts(data)
    })
  }, [])

  return (
    <ScrollArea>
      {
        posts ? (
          posts.map((post) => (
            <div key={post.id}>
              <header className="flex flex-row items-center">
                <Avatar className="size-6">
                  <AvatarImage src="https://github.com/thethoomm.png" />
                  <AvatarFallback>
                    {
                      user.username.slice(0, 2)
                    }
                  </AvatarFallback>
                </Avatar>
                <p className="font-mono text-sm font-semibold">{post.authorUsername}</p>
              </header>
              <main>
                {post.content}
              </main>
              <footer>
                <Button variant={isLiked === post.id ? "default" : "outline"} onClick={() => handleLike(post.id)}>
                  <ArrowUpIcon className="size-4"/>
                  <p>{post.likes}</p>
                </Button>
              </footer>
            </div>
          ))
        ) : (
          <div>No posts</div>
        )
      }
    </ScrollArea>
  )
}