"use client";
import React, { Dispatch, SetStateAction, createContext, useContext, useEffect, useState } from 'react';

type TextColorModeContextType = {
    isWaveEffect: boolean;
    setIsWaveEffect: Dispatch<SetStateAction<boolean>>;
};

const defaultContextValue: TextColorModeContextType = {
    isWaveEffect: false,
    setIsWaveEffect: () => { }, // This now correctly matches the signature
};

const TextColorModeContext = createContext<TextColorModeContextType>(defaultContextValue);

export const useTextColorMode = () => {
    const context = useContext(TextColorModeContext);
    if (!context) {
        throw new Error('useTextColorMode must be used within a TextColorModeProvider');
    }
    return context;
};

export const TextColorModeProvider = ({ children }: { children: React.ReactNode }) => {
    const [isWaveEffect, setIsWaveEffect] = useState(false);

    useEffect(() => {
        if (isWaveEffect) {
            document.documentElement.classList.add('wave-effect');
        } else {
            document.documentElement.classList.remove('wave-effect');
        }

        // Cleanup function to remove class when component unmounts or effect reruns
        return () => {
            document.documentElement.classList.remove('wave-effect');
        };
    }, [isWaveEffect]);

    return (
        <TextColorModeContext.Provider value={{ isWaveEffect, setIsWaveEffect }}>
            {children}
        </TextColorModeContext.Provider>
    );
};
