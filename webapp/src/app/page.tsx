"use client"

import { Navbar } from "@/components/Navbar";
import { useAuth } from "@/contexts/AuthContext";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const router = useRouter()
  const { token } = useAuth()
  
  // useEffect(() => {
  //   if (!token) {
  //     router.push("/login")
  //   }
  // }, [])

  return (
    <main>
      <Navbar />
    </main>
  );
}
