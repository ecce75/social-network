import { useRouter } from 'next/navigation';
import FriendsList from '../friends/FriendsList';
import UserTab from '../friends/UserTab';



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
 
    
    
        const profile = () => {
            router.push('/dashboard/profile/placeholderprofile'); // Redirect to groups page
        };
    

    return (

        <div className="flex-none">
            <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                <div className="w-10 rounded-full">
                <img alt="Tailwind CSS Navbar component" src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                </div>
            </div>
            {/* Dropdown menu */}
            <ul tabIndex={0} className="menu menu-lg dropdown-content mt-5 z-[1] p-3 shadow bg-primary rounded-box w-72 border-2 border-black">

                    <ul className="flex justify-between menu menu-horizontal bg-secondary rounded-box">
                        <li onClick={profile}>
                            <a>
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 20 20" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z" /></svg>
                            </a>
                        </li>
                        <li>
                            <a>
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 20 20" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" /></svg>
                            </a>
                        </li>
                        <li onClick={logout}>
                            <a>
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 20 20" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z" /></svg>
                            </a>
                        </li >
                    </ul>
                               
                <div style={{  backgroundColor: '#4F7942', borderRadius: '8px', height: '50vh', padding: '10px', marginTop: '10px', marginBottom: '20px', overflowY: 'auto' }}>
                
                {/* Friendsist */}
                <ul style={{  display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                {/* "Friend" header */}
                <ul style={{ border: '2px solid #355E3B', borderRadius: '8px',padding: '3px', display: 'flex', alignItems: 'center', backgroundColor: '#355E3B' }}>
                <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" fill="none" viewBox="0 0 20 20" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z" />
                </svg>
                <span>  Friends</span>
                </ul>
                    
                    {/* Map through friends and render each item */}
                    <FriendsList/>
                    <UserTab/>
      

                </ul>
            </div>
            </ul>
            </div>
        </div>
    )
}

export default ProfileIconDM;