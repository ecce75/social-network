"use client"
import MainHeader from '@/components/headers/MainHeader';
import background from '../../../public/assets/background.png';
import { Metadata } from 'next';
import Footer from '@/components/headers/Footer';
import SkeletonFeed from '@/components/feeds/SkeletonFeed';


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