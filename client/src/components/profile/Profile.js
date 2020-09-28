import React from 'react';
import "./Profile.css"
import "../common/Common.css"
import {LOREM} from "../common/Common.js"


export class Profile extends React.Component {

  constructor(props) {
    super(props); 
    this.state = {
      name: "Le Tran Anh Duc",
      shortDescription: "Technology Enthusiast",
      aboutMe: LOREM,
      imageUrl: "res/images/profile.jpg",
      age: 20,
      location: "Remote"
    }
  }

  render() {
    return (
      <div className="Profile" id="Profile">
        <div className="container Profile_Container">
          <text className="container_title">Profile</text>
          <text className="container_subtitle">{this.state.shortDescription}</text>
          <div>
            <div className="Profile_Details">
              <text>About Me</text>
              <text>{this.state.aboutMe}</text>
            </div>
            <img className="round_img Profile_Img" src={this.state.imageUrl} alt="My Profile"/>
            <div className="Profile_Details">
              <text>Details</text>
              <text><b>Name:</b></text>
              <text>{this.state.name}</text>
              <text><b>Age:</b></text>
              <text>{this.state.age}</text>
              <text><b>Location:</b></text>
              <text>{this.state.location}</text>
            </div>
          </div>
        </div>
      </div>
    )
  }
}