import Image from "next/image"
import rastaLionImage from '../public/assets/rasta_lion.png';
import LoginForm from "./Login_Register";

function Header() {
    return (
        <div className="flex justify-between items-center p-4 bg-green-800 text-white">
            <h1>Header</h1>
        {/*left*/}
        <div>
            <Image src={rastaLionImage} alt="Rasta lion" width={50} height={50} />
        </div>
        {/*right*/}
        {/*    Username*/}
        {/*    Notifications*/}
        {/*    Logout*/}
            <LoginForm/>
        </div>
    )
}

export default Header;