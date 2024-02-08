import RastaIcon from "../icons/RastaIcon";
import { useRouter } from 'next/navigation';

function IrieSphereButton (){
    const router = useRouter();
    
    const dashboard = () => {
        router.push('/dashboard');
    };

    return (

        <button onClick={dashboard} className="btn btn-ghost text-3xl font-bold  relative left-20">
        <RastaIcon />
        <span>IrieSphere</span>
        </button>
    )
}

export default IrieSphereButton;