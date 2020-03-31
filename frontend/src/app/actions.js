// redux/actions.js
import { ADD_ENTRY } from './actionTypes'

let nextEntryId = 0
export const addEntry = content => ({
  type: ADD_ENTRY,
  payload: {
    id: ++nextEntryId,
    content
  }
})