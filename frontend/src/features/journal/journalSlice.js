import { createSlice } from "@reduxjs/toolkit";

export const slice = createSlice({
  name: "journal",
  initialState: {
    entries: [
      {
        uuid: "29eabdd4-769b-4550-bc16-3fc432cb95c6",
        title: "Week of Feb 24",
        body:
          "I am, so far, successfully developing an API gateway for online journal"
      },
      {
        uuid: "82e07483-ff16-4052-8fc4-97d043307f18",
        title: "Week of March 22",
        body:
          "Corona Scare is high. Cuomo is being presidential. Trump is impatient with the virus and wants the country to go back to normal by Easter. I have refactored code to be more readable (hopefully). This time for real as I was updating the wrong repo previously."
      }
    ]
  },
});

// export const { addEntry, viewEntry, clearEntry, editEntry } = slice.actions;

// The function below is called a selector and allows us to select a value from
// the state. Selectors can also be defined inline where they're used instead of
// in the slice file. For example: `useSelector((state) => state.counter.value)`
export const selectEntries = state => state.journal.entries;

export default slice.reducer;
