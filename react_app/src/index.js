// =========================================
// Index.js
// ----
// Webpack build entry
// =========================================

// ---- External Dependencies ----
import React from 'react';
import injectTapEventPlugin from 'react-tap-event-plugin';

// ---- Internal Dependencies ----
import App from './components/App';

React.initializeTouchEvents(true);
injectTapEventPlugin();

// ---- Render (entry) ----
if (typeof document !== 'undefined') {
  React.render(<App />, document.getElementById('outlet'));
}

// ==== Module Export ====
export default App;
