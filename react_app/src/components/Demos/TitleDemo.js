// ===========================================
// TitleDemo.js
// ----
// Specific demo for site title
// ===========================================

// ---- External Dependencies ----
import React, { Component } from 'react';

// ---- Internal Dependencies ----
import TextInputDemo from './TextInputDemo';

// ---- Styles ----

// ---- React Class ----
class TitleDemo extends Component {

//  static propTypes = {
//    springConfig: React.PropTypes.arrayOf( React.PropTypes.number )
//  }

  //static defaultProps = {
  //  springConfig: [170, 26]
  //}

  render() {
    return (
      <TextInputDemo
        animKey="shorten"
        headerTxt="URL:"
        explainTxt="Enter a URL to shorten it"
        saveTxt="Go"
        initialInputVal="https://darthvader.darksideclub.com"
        springConfig={ this.props.springConfig } />
    );
  }

};

// ==== Module Export ====
export default TitleDemo;
