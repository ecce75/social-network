import Image from "next/image"
import rastaLionImage from '../../../public/assets/rasta_lion.png';


import { useRouter } from 'next/navigation';

function MainHeader() {
    const router = useRouter();

    const logout = async () => {
        const response = await fetch('http://localhost:8080/api/users/logout', {
            method: 'POST',
            credentials: 'include',
        });

        if (response.ok) {
            // Redirect to login page or show a success message
            router.push('/auth');
        } else {
            // Handle error
            console.error('Logout failed');
        }
    };

    return (
        <div className="flex justify-between items-center p-4 bg-green-800 text-white">
            <div className="flex items-center">
                <Image src={rastaLionImage} priority={true} alt="Rasta lion" width={50} />
                <h1 className="ml-4 font-rasa text-3xl">IrieSphere</h1></div>

            {/*right*/}
            {/*    Username*/}
            {/*    Notifications*/}
            <button onClick={logout}>Logout</button>
        </div>
    )
}


export default MainHeader;