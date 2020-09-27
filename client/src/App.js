import React from 'react';
import './App.css';
import { Header } from './components/header/Header.js'
import { Profile } from './components/profile/Profile.js'
import { Main } from './components/main/Main.js'
import { Experiences } from './components/experiences/Experiences.js'
import { Projects } from './components/projects/Projects.js'
import { Skills } from './components/skills/Skills.js'
import { Contact } from './components/contact/Contact.js'
import { Blogs } from './components/blogs/Blogs.js'

function App() {
  return (
    <div className="App">
      <Header />
      <Main />
      <Profile />
      <Experiences />
      <Projects />
      <Skills /> 
      <Blogs />
      <Contact />
    </div>
  );
}

export default App;