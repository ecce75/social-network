import Image from "next/image"
import rastaLionImage from '../../../public/assets/rasta_lion.png';
function RastaIcon (){

    return (

        <Image src={rastaLionImage} priority={true} alt="Rasta lion" width={55} />
    )
}

export default RastaIcon;