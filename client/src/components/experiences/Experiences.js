import React from 'react';
import './Experiences.css'
import "../common/Common.css"
import "../common/Common.js"
import { LOREM } from '../common/Common.js';

export class Experiences extends React.Component {

  constructor(props) {
    super(props); 
    this.state = {

    }
  }

  render() {
    return (
      <div className="Experiences" id="Experiences">
        <div className="container">
          <text className="container_title">Experiences</text>
          <text className="container_subtitle">Some deep quote about experiences</text>
          <Experience />
        </div>
      </div>
    )
  }
}

// TODO: This can just be a single function
class Experience extends React.Component {

  constructor(props) {
    super(props); 
    this.state = {
        organization:"Organization Name",
        organizationUrl:"Organization.url",
        role:"Software Developer",
        description:LOREM,
        type:"Fulltime",
        time:"Sometime - Present"
    }
  }

  render() {
    return (
      <div className="Experience">
        <div className="Experience_Main">
          <text className="Experience_Org">{this.state.organization}</text>
          <text className="Experience_Time">{this.state.time}</text>
        </div>
        <div className="Experience_Details">
          <text className="Experience_Title">{this.state.type} - {this.state.role}</text>
          <text>{this.state.description}</text>
          <text>{this.organizationUrl}</text>
        </div>
      </div>
    )
  }
}