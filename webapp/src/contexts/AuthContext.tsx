"use client"
import { ReactNode, createContext, useContext, useState } from "react"
import { jwtDecode } from 'jwt-decode'

interface AuthContextType {
  token: string | null | undefined;
  setToken: React.Dispatch<React.SetStateAction<string | null >>
}

interface Props {
  children: ReactNode
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: Props) {
  const localToken = localStorage.getItem("token")
  const [token, setToken] = useState<string | null >(localToken)

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