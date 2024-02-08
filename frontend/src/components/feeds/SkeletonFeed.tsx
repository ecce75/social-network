import React from 'react';
import BigBoxSkeleton from '../skeletons.tsx/BigBoxSkeleton';

const SkeletonFeed: React.FC = () => {
    return (
        /* Change % for post feed width*/
        <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', maxHeight: '110vh', overflowY: 'auto' }}>
            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                <BigBoxSkeleton/>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                <BigBoxSkeleton/>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                <BigBoxSkeleton/>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                <BigBoxSkeleton/>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                <BigBoxSkeleton/>
            </div>
        </section>
    );
};

export default SkeletonFeed;
