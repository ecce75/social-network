"use client";

import NotificationButton from "../buttons/NotificationButton";
import ProfileIconDM from "../buttons/ProfileIconDropdown";
import LeftNavBar from "../leftnavbar/LeftNavBar";
import IrieSphereButton from "../buttons/IriesphereButton";
import React, {useEffect, useState} from "react";

function MainHeader() {
    const [friendsListToggle, setFriendsListToggle] = React.useState<boolean>(false)
    return (
        <div className="navbar bg-primary h-">
            <div className="flex-1">
                <LeftNavBar />
                <IrieSphereButton />
            </div>
            <div className="navbar-end">
                {/* <ToggleWaveEffectButton /> */}
                <div>
                    <NotificationButton
                        setFriendsListToggle={setFriendsListToggle}
                    />
                </div>
                <div>
                    <ProfileIconDM
                        key = {friendsListToggle.toString()}
                        friendsListToggle={friendsListToggle}
                    />
                </div>

            </div>
        </div>

    )
}


export default MainHeader;