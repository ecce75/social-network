import Image from "next/image"
import LoginForm from "../auth/Login";


function LoginHeader() {
    return (
        <div className="flex justify-between items-center p-4 bg-primary text-white">
            <div className="flex items-center">
                <Image src={"/assets/rasta_lion.png"} priority={true} alt="Rasta lion" width={50} height={50} />
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