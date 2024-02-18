import Footer from '@/components/headers/Footer';
import background from '../../../public/assets/background.png';

export default function AuthLayout({
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
        {children}
        <Footer/>
        </div>
      
        
    )
  }