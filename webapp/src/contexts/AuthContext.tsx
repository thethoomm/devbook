"use client"
import { ReactNode, createContext, useContext, useState } from "react"

interface AuthContextType {
  token: string | null | undefined;
  setToken: React.Dispatch<React.SetStateAction<string | null | undefined>>
}

interface Props {
  children: ReactNode
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: Props) {
  const [token, setToken] = useState<string | null >()

  return (
    <AuthContext.Provider value={{ token, setToken }}>
      { children }
    </AuthContext.Provider>
  )
} 

export function useAuth() {
  const context = useContext(AuthContext)
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider')
 }
 return context
}