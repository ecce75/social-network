import Image from 'next/image'
import Head from 'next/head'
import Header from '../components/Header'
export default function Home() {
  return (

      <div >
          <Header />
        <Head>
          <title>IrieSphere</title>
        </Head>
        <h1>Rastafari social-network</h1>
      </div>
  )
}
