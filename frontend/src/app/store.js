import { configureStore } from '@reduxjs/toolkit';
import counterReducer from '../features/counter/counterSlice';
import journalReducer from '../features/journal/journalSlice';

export default configureStore({
  reducer: {
    counter: counterReducer,
    journal: journalReducer,
  },
});
