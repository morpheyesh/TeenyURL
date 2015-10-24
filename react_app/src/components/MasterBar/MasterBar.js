// =========================================
// MasterBar.js
// ----
// Fake master bar for sticky header
// =========================================

// ---- External Dependencies ----
import React, { Component } from 'react';

// ---- Internal Dependencies ----

// ---- Styles ----
import styles from './MasterBar.css';

// ---- React Class ----
class MasterBar extends Component {

  render() {
    return (
      <div className={ styles.bar } >
        <h2 className={ styles.titz }>TeenyUrl</h2>
      </div>
    );
  }

};

// ==== Module Export ====
export default MasterBar;
