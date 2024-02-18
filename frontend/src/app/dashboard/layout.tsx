"use client"
import MainHeader from '@/components/headers/MainHeader';
import background from '../../../public/assets/background.png';
import Footer from '@/components/headers/Footer';





export default function DashboardLayout({
    children,
  }: {
    children: React.ReactNode
  }) {
    return (
        <div style={{
          backgroundImage: `url("${background.src}")`,
          backgroundSize: 'cover',
          flex: 1,

      }}>
        
        <MainHeader/>
        {children}
        
        <Footer/>
        </div>
      
        
    )
  }

