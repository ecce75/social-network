import NotificationButton from "../buttons/NotificationButton";
import ProfileIconDM from "../buttons/ProfileIconDropdown";
import LeftNavBar from "../leftnavbar/LeftNavBar";
import IrieSphereButton from "../buttons/IriesphereButton";
import { ToggleWaveEffectButton } from "@/components/buttons/ToggleColorModeButton";

function MainHeader() {
    return (
        <div className="navbar bg-primary h-">
            <div className="flex-1">
                <LeftNavBar />
                <IrieSphereButton />
            </div>
            <div className="navbar-end">
                {/* <ToggleWaveEffectButton /> */}
                <div>
                    <NotificationButton />
                </div>
                <div>
                    <ProfileIconDM />
                </div>

            </div>
        </div>

    )
}


export default MainHeader;