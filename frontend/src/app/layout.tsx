
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import '../../styles/globals.css'
import background from '../../public/assets/background.png';
import dynamic from 'next/dynamic';

const inter = Inter({ subsets: ['latin'] })
const MainHeader = dynamic(() => import('@/components/headers/MainHeader'), { ssr: false });
import AuthLayout from './auth/layout';

export const metadata: Metadata = {
  title: 'Iriesphere',
  description: 'Rastafari social-network',
}



export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      
      <body className={inter.className}>
        {/* <MainHeader /> */}
        <div style={{
          backgroundImage: `url("${background.src}")`,
          backgroundSize: 'cover',
          flex: 1,
          minHeight: '100vh',
      }}>{children}</div>
            
            </body>
       
    </html>
  )
}
