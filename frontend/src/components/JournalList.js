import React from "react";
import Entry from "./Entry";
import { connect } from "react-redux";
import { getEntries } from "../redux/selectors";


const JournalList = ({ entries }) => (
  <ul className="journal-list">
    {entries && entries.length
      ? entries.map((entry, index) => {
          return <Entry key={`entry-${entry.id}`} entry={entry} />;
        })
      : "No Entries, yay!"}
  </ul>
);
  
export default connect(state => ({ entries: getEntries(state) }))(JournalList);