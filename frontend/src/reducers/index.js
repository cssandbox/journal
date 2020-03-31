import { combineReducers } from 'redux'
import journalReducer from '../features/journal/journalSlice'

export default combineReducers({
  journal: journalReducer
})