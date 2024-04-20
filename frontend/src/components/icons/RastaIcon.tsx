import Image from "next/image"

function RastaIcon() {
    return (
        <Image src={"/assets/rasta_lion.png"} priority={true} alt="Rasta lion" width={55} />
    )
}

export default RastaIcon;