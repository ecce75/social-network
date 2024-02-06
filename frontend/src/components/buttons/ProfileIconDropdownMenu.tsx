import { useRouter } from 'next/navigation';



function ProfileIconDM (){
    
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

        <div className="flex-none">
            <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                <div className="w-10 rounded-full">
                <img alt="Tailwind CSS Navbar component" src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                </div>
            </div>
            {/* Dropdown menu */}
            <ul tabIndex={0} className="menu menu-lg dropdown-content mt-3 z-[1] p-2 shadow bg-primary rounded-box w-52 border-2 border-black">
                <li className="border-b border-black rounded-box"><a>Profile</a></li>
                <li className="border-b border-black rounded-box"><a>Settings</a></li>
                <li onClick={logout} className="border-b border-black rounded-box"><a>Logout</a></li>
            </ul>
            </div>
        </div>
    )
}

export default ProfileIconDM;