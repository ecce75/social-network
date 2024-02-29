import React from 'react';
import EventsButton from '../buttons/EventsButton';
import NavGroupsButton from '../buttons/NavGroupsButton';

function LeftNavBar() {
  return (
    <div className="drawer" style={{ zIndex: 9999 }}>
      <input id="my-drawer" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        <label htmlFor="my-drawer" className="btn bg-primary swap swap-rotate border-b border-white round">
          <input type="checkbox" />
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 512 512">
  <path fill="white" d="M64,384H448V341.33H64Zm0-106.67H448V234.67H64ZM64,128v42.67H448V128Z"/>
</svg>
        </label>
      </div> 
      <div className="drawer-side">
        <label htmlFor="my-drawer" aria-label="close sidebar" className="drawer-overlay"></label>
        <ul className="menu p-4 w-80 min-h-full bg-primary text-base-content">
          {/* Sidebar buttons here */}
          <li><NavGroupsButton /></li>
          <li><EventsButton /></li>
        </ul>
      </div>
    </div>
  );
}

export default LeftNavBar;
