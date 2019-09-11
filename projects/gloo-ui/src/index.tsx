import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './fontFace.css';
import { GlooIApp } from './GlooIApp';
import * as serviceWorker from './serviceWorker';
import { globalStore } from './store';

ReactDOM.render(
  <Provider store={globalStore}>
    <GlooIApp />
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
