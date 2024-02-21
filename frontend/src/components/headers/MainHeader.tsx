import NotificationButton from "../buttons/NotificationButton";
import ProfileIconDM from "../buttons/ProfileIconDropdown";
import LeftNavBar from "../leftnavbar/LeftNavBar";
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
            
            <ProfileIconDM />
        </div>
        </div>
    
    )
}


export default MainHeader;