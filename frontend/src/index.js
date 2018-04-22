import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import Login from './components/login'

import registerServiceWorker from './registerServiceWorker';

//ReactDOM.render(<App />, document.getElementById('1root'));
ReactDOM.render(<Login />, document.getElementById('root-container'));

registerServiceWorker();
