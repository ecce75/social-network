import Image from "next/image"
import rastaLionImage from '../../../public/assets/rasta_lion.png';
import FriendsButton from "../buttons/FriendsButton";
import NotificationButton from "../buttons/NotificationButton";
import ProfileIconDM from "../buttons/ProfileIconDropdownMenu";
import RastaIcon from "../icons/RastaIcon";
import LeftNavBar from "../leftNavbar/LeftNavBar";
import IrieSphereButton from "../buttons/IriesphereButton";

function MainHeader() {
        {/* Logout logic at buttons/ProfileIconDropdownMenu */}
    return (
        
        <div className="navbar bg-primary h-">
        <div className="flex-1">
            <LeftNavBar />
            <IrieSphereButton />
        </div>
        <div className="navbar-end">
            <NotificationButton />
            <FriendsButton />
            <ProfileIconDM />
        </div>
        </div>
    
    )
}


export default MainHeader;