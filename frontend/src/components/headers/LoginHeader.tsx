import Image from "next/image"
import rastaLionImage from '../../../public/assets/rasta_lion.png';
import { LoginForm } from "../auth/LoginRegister";


function LoginHeader() {
    return (
        <div className="flex justify-between items-center p-4 bg-primary text-white">
            <div className="flex items-center">
                <Image src={rastaLionImage} priority={true} alt="Rasta lion" width={50} height={50} />
                <h1 className="ml-4 font-rasa text-3xl">IrieSphere</h1></div>
            {/*right*/}
            {/*    Username*/}
            {/*    Notifications*/}
            {/*    Logout*/}
            <LoginForm />
        </div>
    )
}

export default LoginHeader;