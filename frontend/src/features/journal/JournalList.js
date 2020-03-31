import React from "react";
import Entry from "./Entry";
import { connect } from "react-redux";

const JournalList = ({ entries }) => (
  <ul className="journal-list">
    {entries && entries.length
      ? entries.map((entry, index) => {
          return <Entry key={`entry-${entry.id}`} entry={entry} />;
        })
      : "No Entries, yay!"}
  </ul>
);
  

const mapStateToProps = state => ({
    entries: state.journal
  })
  
export default connect(
    mapStateToProps,
    null
  )(JournalList)

