import { createSlice } from '@reduxjs/toolkit'

let nextEntryId = 0

const journalSlice = createSlice({
  name: 'journal',
  initialState: [],
  reducers: {
    addEntry: {
      reducer(state, action) {
        const { id, title } = action.payload
        state.push({ id, title })
      },
      prepare(title) {
        return { payload: { title, id: nextEntryId++ } }
      }
    }
  }
})

export const { addEntry } = journalSlice.actions

export default journalSlice.reducer