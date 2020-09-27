import React from 'react';
import './App.css';
import { Header } from './components/Header.js'
import { Experiences } from './components/Experiences.js'
import { Projects } from './components/Projects.js'
import { Skills } from './components/Skills.js'
import { Contact } from './components/Contact.js'
import { Blogs } from './components/Blogs.js'

function App() {
  return (
    <div className="App">
      <Header />
      <Experiences m/>
      <Projects />
      <Skills /> 
      <Blogs />
      <Contact />
    </div>
  );
}

export default App;
