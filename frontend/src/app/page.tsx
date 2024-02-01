import Image from 'next/image'
import Head from 'next/head'
import Header from '../components/Header'
import {RegisterForm}from '../components/Login_Register'
import background from '../public/assets/background.png';
import rastaLionImage from '../public/assets/rasta_lion.png';

export default function Home() {
  return (

      <div className="flex flex-col h-screen">
          <Header/>
          <Head>
              <title>IrieSphere</title>
          </Head>
        <div style={{
            backgroundImage: `url("${background.src}")`,
            backgroundSize: 'cover',
            flex: 1,
            
        }}>
            <RegisterForm/>
        </div>
      </div>

  )
}
