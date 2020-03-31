import { createSlice } from "@reduxjs/toolkit";

let nextEntryId = 0;

const journalSlice = createSlice({
  name: "journal",
  initialState: [],
  reducers: {
    addEntry: {
      reducer(state, action) {
        const { uuid, title, body } = action.payload;
        state.push({ uuid, title, body });
      },
      prepare(title, body) {
        return { payload: { title, body, uuid: nextEntryId++ } };
      }
    }
  }
});

export const { addEntry } = journalSlice.actions;

export default journalSlice.reducer;
