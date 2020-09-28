import React from 'react';
import "./Profile.css"
import "../common/Common.css"

const lorem = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

export class Profile extends React.Component {

  constructor(props) {
    super(props); 
    this.state = {
      name: "Le Tran Anh Duc",
      shortDescription: "Technology Enthusiast",
      aboutMe: lorem,
      imageUrl: "res/images/profile.jpg",
      age: 20,
      location: "Remote"
    }
  }

  render() {
    return (
      <div className="Profile" id="Profile">
        <div className="container Profile_Container">
          <text>Profile</text>
          <text>{this.state.shortDescription}</text>
          <div className="Profile_Detail">
            <text>About Me</text>
            <text>{this.state.aboutMe}</text>
          </div>
          <img className="round_img Profile_Img" src={this.state.imageUrl}/>
          <div className="Profile_Detail">
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
    )
  }
}