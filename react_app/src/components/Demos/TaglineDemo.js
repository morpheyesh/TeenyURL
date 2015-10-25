// ===========================================
// TaglineDemo.js
// ----
// Specific demo for site title
// ===========================================

// ---- External Dependencies ----
import React, { Component } from 'react';

// ---- Internal Dependencies ----
import TextInputDemo from './TextInputDemo';

// ---- Styles ----

// ---- React Class ----
class TaglineDemo extends Component {

  static propTypes = {
    springConfig: React.PropTypes.arrayOf( React.PropTypes.number )
  }

  static defaultProps = {
    springConfig: [170, 26]
  }

  render() {
    return (
      <TextInputDemo
        animKey="tagline"
        headerTxt="TeenyURL:"
        explainTxt="Enter a teenyurl to lengthen it"
        editTxt=""
        saveTxt="Go"
        initialInputVal="teenyurl.co/xF1trL"
        springConfig={ this.props.springConfig } />
    );
  }

};

// ==== Module Export ====
export default TaglineDemo;
