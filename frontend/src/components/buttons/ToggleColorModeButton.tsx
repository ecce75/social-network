import { useTextColorMode } from '@/hooks/TextColorModeContext';

export const ToggleWaveEffectButton = () => {
    const { isWaveEffect, setIsWaveEffect } = useTextColorMode() as { isWaveEffect: boolean, setIsWaveEffect: (value: boolean) => void };

    const toggleWaveEffect = () => {
        setIsWaveEffect(!isWaveEffect);
    };

    const buttonStyle = {
        width: 'fit-content',
        height: 'auto',
        border: '2px red',
        marginRight: '5%'
    };

    return (
        <div>
            {isWaveEffect ? <button style={buttonStyle} onClick={toggleWaveEffect}>WAKE UP!</button> : <button className='createdAtWave' onClick={toggleWaveEffect} style={buttonStyle}>Ü.L.E</button>}
        </div>
            
    );
};
