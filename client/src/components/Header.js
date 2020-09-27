import React from 'react';
import "./Header.css"

export function Header() {
    return (
      <div className="Header">
        <text className="Logo" href="#Main">DL</text>
        <div className="Menu">
          <text className="MenuItem" href="#AboutMe">About Me</text>
          <text className="MenuItem" href="#Experiences">Experiences</text>
          <text className="MenuItem" href="#Skills">Skills</text>
          <text className="MenuItem" href="#Projects">Projects</text>
          <text className="MenuItem" href="#Blogs">Blogs</text>
          <text className="MenuItem" href="#Contact">Contact</text>
        </div>
      </div>
    )
  }