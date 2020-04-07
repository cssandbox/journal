import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import { Provider } from 'react-redux';
import * as serviceWorker from './serviceWorker';
import rootReducer from "./reducers";
import { configureStore } from "@reduxjs/toolkit";
import thunkMiddleware from 'redux-thunk'
import fetchResource from './services/api'
import { addEntry } from './features/journal/journalSlice'

function getEntries() {
  return fetchResource('entries');
}

function populateFromServer() {
  return function(dispatch) {
    return getEntries().then(entriesData => {
      entriesData.map(entry => {
        dispatch(addEntry(entry.title, entry.body))
      })
    })
    .catch(error => {
      console.log(error)
    })
  }
}

const store = configureStore({
  reducer: rootReducer,
  middleware: [thunkMiddleware],
});

store.dispatch(populateFromServer())

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
