"use client"

import Head from 'next/head'
import Header from '../../components/headers/LoginHeader'
import { RegisterForm } from '../../components/auth/LoginRegister'
import background from '../../../public/assets/background.png';


export default function Auth()  {
    return (

        <div className="flex flex-col h-screen">
            <Header />
            <Head>
                <title>IrieSphere</title>
            </Head>
            <div><RegisterForm />
            </div>
            
                
            </div>

    )
}